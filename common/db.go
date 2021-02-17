package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func InitDb(){
	drivername := viper.GetString("db.drivername")
	var args string

	//选择链接的数据库
	switch drivername {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			viper.GetString("db.dbuser"),
			viper.GetString("db.dbpwd"),
			viper.GetString("db.hostname"),
			viper.GetString("db.port"),
			viper.GetString("db.dbname"),
			viper.GetString("db.charset"),
		)

	default:
		panic("不支持的数据库")
	}

	db, err := gorm.Open(drivername, args)
	if err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("db.lifetime")) * time.Second)
	db.DB().SetMaxOpenConns(viper.GetInt("db.maxconn"))
	db.DB().SetMaxIdleConns(viper.GetInt("db.idleconn"))
	Db = db
	return
}

func GetDb() *gorm.DB {
	return Db
}
