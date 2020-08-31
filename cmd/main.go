package main

import (
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "swaggo_sample/docs"
	h "swaggo_sample/internal/handler"
)

// @title gin-swagger sample
// @version 0.1
func main() {
	var Server = os.Getenv("SERVER")
	var Port = os.Getenv("PORT")

	var Router *gin.Engine
	Router = gin.Default()

	Router.GET("/instances", h.InstanceHandler)
	Router.GET("/images", h.ImageHandler)

	url := ginSwagger.URL("http://" + Server + ":" + Port + "/swagger/doc.json")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	Router.Run(":" + Port)
}
