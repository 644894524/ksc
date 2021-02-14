package router

import (
	"github.com/gin-gonic/gin"
	"ksc/controller/article"
)

func newsModule(app *gin.Engine) *gin.Engine {
	us := app.Group("article")
	{
		us.GET("/List", article.List)
	}
	return app
}

