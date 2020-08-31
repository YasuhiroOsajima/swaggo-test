package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
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
	bytes, err := ioutil.ReadFile("internal/handler/instances.json")
	if err != nil {
		log.Fatal(err)
	}

	var instances []Instance
	if err := json.Unmarshal(bytes, &instances); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, instances)
}
