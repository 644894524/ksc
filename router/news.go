package router

import (
	"github.com/gin-gonic/gin"
	"ksc/controller/news"
)

func newsModule(app *gin.Engine) *gin.Engine {
	us := app.Group("news")
	{
		us.GET("/List", news.List)
	}
	return app
}

