package main

import (
	"github.com/derekAHua/goLib"
	"github.com/derekAHua/goLib/server/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goWeb/conf"
	"goWeb/router"
)

func main() {
	engine := gin.New()

	goLib.Bootstraps(engine)

	conf.Init()
	defer conf.Clear()

	// go run main.go goWeb
	var rootCmd = &cobra.Command{
		Use:   "goWeb",
		Short: "goWeb application",
		Run: func(cmd *cobra.Command, args []string) {
			httpServer(engine)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

// web server: http web 服务，容器里监听特定端口(:8080)实现web请求
func httpServer(engine *gin.Engine) {
	// 初始化http服务路由
	router.Http(engine)
	// MQ 消费回调路由
	//router.MQ(engine)

	// 启动web server
	if err := http.Start(engine, conf.Config.Server); err != nil {
		panic(err.Error())
	}
}
