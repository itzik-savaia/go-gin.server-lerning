package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/entity"
	"main.go/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (con *controller) FindAll() []entity.Video {
	return con.service.FindAll()
}

func (con *controller) Save(c *gin.Context) entity.Video {
	var video entity.Video
	c.BindJSON(&video)
	con.service.Save(video)
	return video
}
