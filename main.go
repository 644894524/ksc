package main

import (
	"ksc/router"
	"ksc/common"
)

func main() {
	common.InitViper()
	common.InitDb()

	//gin （路由）
	router.RoutersInit()
}
