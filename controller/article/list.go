package article

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"ksc/controller"
	"ksc/model/article"
)

func List(c *gin.Context){
	list := article.List(0)
	info, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	//c.JSON(200, info)
	controller.Success(c, gin.H{
		"info":info,
	}, "success")
}
