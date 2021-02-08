package router

import (
	"github.com/gin-gonic/gin"
	"ksc/controller/user"
)

func userModule(app *gin.Engine) *gin.Engine {
	us := app.Group("user")
	{
		us.GET("/UserInfo", user.UserInfo)
	}
	return app
}

