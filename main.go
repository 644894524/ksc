package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"ksc/common"
	"ksc/router"
	"ksc/script"
)

func main() {
	common.InitViper()
	common.InitDb()

	var mode string
	flag.StringVar(&mode, "mode", "web", "task name")
	flag.Parse()
	initMode(mode)
}

func initMode(d string){
	if d == "web" {
		router.RoutersInit()
	}else {
		script.Task()
	}
}
