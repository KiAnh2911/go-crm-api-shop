package initialize

import (
	c "github.com/KiAnh2911/go-crm-api-shop/internal/controller"
	m "github.com/KiAnh2911/go-crm-api-shop/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(m.AuthenMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		// v1.POST("/ping", Pong)
		// v1.PUT("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}
