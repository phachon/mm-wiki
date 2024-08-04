package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/router"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化命令行
	global.InitFlagEnvVar()
	// 初始化配置
	config.Init()
	// 初始化日志
	logger.Init()
	defer logger.Sync()
	// 初始化路由和中间件
	router.Init()
	// 初始化DB
	dao.InitDB()
	// 设置 gin 框架允许环境
	gin.SetMode(config.GetAppConf().Global.GinMode)

	// 启动 server
	s := &http.Server{
		Addr:           config.GetServerAddr(),
		Handler:        global.GinEngine,
		ReadTimeout:    config.GetReadTimeout(),
		WriteTimeout:   config.GetWriteTimeout(),
		MaxHeaderBytes: 1 << 20,
	}
	logger.Infof("[mm-wiki] Listen addr=%+v", config.GetServerAddr())
	err := s.ListenAndServe()
	if err != nil {
		logger.Fatalf("[mm-wiki] err:%s", err)
	}
}
