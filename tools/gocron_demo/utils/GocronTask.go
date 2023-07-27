package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func task() {
	log.Info("111111111111111111111111111")
	fmt.Println("task函数执行一次", time.Now().Format("2006-01-02 15:04:05"))

}

func taskWithParams(a int, b string) {
	log.Info("2222222222222222222222222222")
	fmt.Println(a, b)
}

func GocronDemo() {
	fmt.Println("开始运行定时任务!!!")

	// 申请一个调度器
	s1 := gocron.NewScheduler()

	// 每三秒执行一次
	s1.Every(1).Second().Do(taskWithParams, 1, "hello")
	s1.Every(2).Seconds().Do(task)

	<-s1.Start()

}
