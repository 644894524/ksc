package router

import (
	"github.com/gin-gonic/gin"
	"ksc/actions"
	"github.com/gin-contrib/cors"
	"github.com/dvwright/xss-mw"
	"github.com/spf13/viper"
	"github.com/gin-contrib/pprof"
)

func RoutersInit() {

	//全局配置文件
	viper.SetConfigFile("./config/conf.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()

	//性能分析工具
	pprof.Register(r, "pprof")

	//开启CORS 避免跨越请求（不使用JSONP的原因，JSONP只支持GET请求）
	config := cors.DefaultConfig()
	site := viper.GetString("site.domain")
	config.AllowOrigins = []string{site}
	r.Use(cors.New(config))

	//XSS过滤
	var xssMdlwr xss.XssMw
	r.Use(xssMdlwr.RemoveXss())

	//用户相关模块
	us := r.Group("user")
	{
		us.GET("/UserInfo", actions.UserInfo)
	}

	//监听端口默认为8080
	r.Run(":8000")
}


