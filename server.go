package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
	"main.go/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.Save(c))
	})

	server.Run(":8080")
}
