package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"swaggo_sample/internal/repository"
)

type Instance struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// InstanceHandler is to get instances info
// @Summary Get instances info
// @Description A list of instances
// @Accept  plain
// @Produce  json
// @Success 200
// @Router /instances [get]
func InstanceHandler(c *gin.Context) {
	db := repository.NewInstanceRepository()
	instancelist, err := db.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	var instances []Instance
	for _, i := range instancelist {
		var ins Instance
		ins.Uuid = i.Uuid
		ins.Name = i.Name
		ins.Owner = i.Owner
		instances = append(instances, ins)
	}

	c.JSON(200, instances)
}
