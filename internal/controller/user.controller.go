package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tantrungse/go-ecommerce-backend-api/internal/service"
	"github.com/tantrungse/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	// response.SuccessResponse(c, 20001, []string{"tipjs", "cr7", "m10"})
	response.ErrorResponse(c, 20003)
}
