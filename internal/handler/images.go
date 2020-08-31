package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Image struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// ImageHandler is to get images info
// @Summary Get images info
// @Description A list of images
// @Accept  plain
// @Produce  json
// @Success 200
// @Router /images [get]
func ImageHandler(c *gin.Context) {
	bytes, err := ioutil.ReadFile("internal/handler/images.json")
	if err != nil {
		log.Fatal(err)
	}

	var images []Image
	if err := json.Unmarshal(bytes, &images); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, images)
}
