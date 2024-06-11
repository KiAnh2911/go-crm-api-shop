package controller

import (
	"github.com/KiAnh2911/go-crm-api-shop/internal/services"
	"github.com/KiAnh2911/go-crm-api-shop/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// uc user controller
// us user service
// ur user repo
// controller => service => repo => model => dbs

func (uc *UserController) GetUserByID(c *gin.Context) {
	// response.SuccessRespone(c, response.ErrorCodeSuccess, []string{"sadasdjasdjk"})
	response.ErrorRespone(c, 40003, "có lỗi xảy ra")
}
