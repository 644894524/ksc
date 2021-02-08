package router

import (
	"github.com/gin-gonic/gin"
	"ksc/controller"
)

func userModule(app *gin.Engine) *gin.Engine {
	us := app.Group("user")
	{
		us.GET("/UserInfo", controller.UserInfo)
	}
	return app
}

