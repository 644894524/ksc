package article

import (
	"github.com/gin-gonic/gin"
	"ksc/controller"
	"ksc/model/article"
)

func List(c *gin.Context){
	list := article.List(0)
	//c.JSON(200, info)
	controller.Success(c, gin.H{
		"list":list,
	}, "success")
}
