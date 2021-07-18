package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-template/dao/mysql"
	"gin-template/dao/redis"
	"gin-template/logger"
	"gin-template/routers"
	"gin-template/settings"

	"go.uber.org/zap"
)

func main() {
	// 读取配置
	err := settings.Init()
	if err != nil {
		fmt.Printf("配置初始化失败 %v\n", err)
		return
	}

	// 初始化日志
	err = logger.Init(&settings.Config.Log)
	if err != nil {
		fmt.Printf("Logger初始化失败 %v\n", err)
		return
	}
	lg := zap.L()
	defer lg.Sync()

	// 初始化MySQL链接
	err = mysql.Init(&settings.Config.Mysql)
	if err != nil {
		fmt.Printf("Mysql初始化失败 %v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化Redis链接
	err = redis.Init(&settings.Config.Redis)
	if err != nil {
		fmt.Printf("Redis初始化失败 %v\n", err)
		return
	}
	defer redis.Close()
	
	// 注册路由
	r := routers.Setup()

	// 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.App.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lg.Error("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	lg.Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		lg.Fatal("Server Shutdown: ", zap.Error(err))
	}

	lg.Info("Server exiting")
}
