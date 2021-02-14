package news

import (
	"github.com/gin-gonic/gin"
	"ksc/controller"
	"ksc/model/article"
)

func List(c *gin.Context){

	article.List(0)

	//c.JSON(200, info)
	controller.Success(c, gin.H{
		"info":"list",
	}, "success")
}
