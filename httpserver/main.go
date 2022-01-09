package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// 读取当前系统的环境变量中的 VERSION 配置
func getEnvVersion() string {
	ver := os.Getenv("VERSION")
	return ver
}

func main() {
	router := gin.Default()

	// 接收客户端 request
	router.GET("/healthz", func(c *gin.Context) {
		glog.V(2).Info("Start http server...")
		reqIP := c.ClientIP()
		if reqIP == "::1" {
			reqIP = "127.0.0.1"
		}
		// Server 端记录访问日志包括客户端 IP，HTTP 返回码
		glog.V(2).Info("ip:", reqIP)
		glog.V(2).Info("status:", http.StatusOK)
		fmt.Printf("ip:%s\n", reqIP)
		fmt.Printf("status:%d\n", http.StatusOK)
		c.Writer.Header().Add("VERSION", getEnvVersion())
		c.String(http.StatusOK, getEnvVersion())
		glog.V(2).Info("end http server...")
	})

	router.Run(":8081")
}
