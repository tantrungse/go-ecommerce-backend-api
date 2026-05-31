package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/tantrungse/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1 := r.Group("v1/2026")
	{
		// Define a simple GET endpoint
		v1.GET("/ping", c.NewPongController().Pong) // v1/2026/ping
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}
