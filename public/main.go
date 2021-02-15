package main

import (
	"ksc/router"
	"ksc/common"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	common.InitViper()
	common.InitDb()

	//gin （路由）
	router.RoutersInit()
}
