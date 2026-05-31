package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Thomas Vo")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...pong " + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "Thomas"},
	})
}