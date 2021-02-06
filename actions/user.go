package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

type User struct {
	Uid  	 int 	`json:"uid"`
	Name 	 string `json:"name"`
	Age  	 int 	`json:"age"`
	ClassId  int 	`json:"class_id"`
	SchoolId int 	`json:"school_id"`
	Phone    string `json:"phone"`
	Domain   string `json:"domain"`
}

//登录
func UserInfo(c *gin.Context) {
	uid, err := strconv.Atoi(c.DefaultQuery("uid", "0"))

	if err != nil {
		c.String(500, "错误")
	}

	domain := viper.GetString("site.domain")
	info := User{Uid: uid, Name: "zhangsan", Age: 18, ClassId: 20, SchoolId: 20, Phone: "13621075519", Domain: domain}
	if err != nil {
		c.String(500, "错误")
		return
	}

	c.JSON(200, info)
	c.String()
}
