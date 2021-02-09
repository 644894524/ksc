package router

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	//设置gin模式
	if viper.GetString("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}else{
		gin.SetMode(gin.DebugMode)
	}

	//启动带有中间件的路由：Logger、Recovery 中间件
	r := gin.Default()

	//不启用中间件：
	//r := gin.New()

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
	initMoudle(r)

	//监听端口默认为8080
	r.Run(":8000")
}

//初始化路由
func initMoudle(app *gin.Engine){
	userModule(app)
}


