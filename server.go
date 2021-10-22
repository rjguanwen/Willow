package main

import (
	"fmt"
	"github.com/rjguanwen/willow/service"
	"net/http"
	"runtime"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/rjguanwen/willow/public"
)

// 初始化
func init() {
	// 设置使用的 CPU 数量
	runtime.GOMAXPROCS(4)
	// 初始化日志配置
	setupLogger()
}

// 设置日志功能配置
func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("config/seelog.xml")
	if err != nil {
		return
	}
	log.ReplaceLogger(logger)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping ok")
	})

	r.GET("/getMembers", func(c *gin.Context) {
		log.Info("/getMembers ----- ")
		index := c.Query("k")
		args := c.Query("v")
		log.Debugf("==>参数 index:%s", index)
		log.Debugf("==>参数 args:%s", args)

		if index == "" {
			c.String(http.StatusOK, "get error，index is null")
		}else if args == "" {
			c.String(http.StatusOK, "get error，args is null")
		}else {
			result, err := service.QueryMembersByTags(index, args)
			if err != nil {
				c.String(http.StatusNotFound, fmt.Sprint(err))
			}
			rstJson, err := json.Marshal(result)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprint(err))
			}
			c.String(http.StatusOK, string(rstJson))
		}
	})

	return r
}

func main() {
	log.Info("Server Start ...")
	log.Info("Config Loading ...")
	baseConfig, _ := public.InitBaseConfig("dev")
	//fmt.Println(baseConfig.Port)

	//log.Info("DB Init ...")
	// 设置数据文件路径
	//service.SetTagsFile(baseConfig.MemberTagsFile)
	//// 初始化数据库
	//service.InitDB()
	log.Info("Server Ready ...")
	r := setupRouter() // Listen and Server in 0.0.0.0:8080
	r.Run(":"+baseConfig.Port)
}
