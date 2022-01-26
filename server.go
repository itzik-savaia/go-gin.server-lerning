package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"main.go/controller"
	"main.go/middlewares"
	"main.go/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, VideoController.Save(c))
	})

	server.Run(":8080")
}
