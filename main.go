package main

import (
	"ksc/common"
	"ksc/router"
	"ksc/util"
	"os"
)

func main(){
	currDir, _ := os.Getwd()
	appDir := util.StringBuilder(currDir, "/")
	common.InitViper(appDir)
	common.InitDb()
	router.RoutersInit(appDir)
}