package article

import (
	"github.com/gin-gonic/gin"
	"ksc/controller"
	"ksc/model"
)

func List(c *gin.Context){
	article := new(model.Article)
	list := article.List(0, 20)
	controller.Success(c, gin.H{
		"list":list,
	}, "success")
}
