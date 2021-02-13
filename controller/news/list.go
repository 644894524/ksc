package news

import (
	"github.com/gin-gonic/gin"
	"ksc/controller"
)

func List(c *gin.Context){
	//c.JSON(200, info)
	controller.Success(c, gin.H{
		"info":"list",
	}, "success")
}
