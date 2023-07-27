package utils

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func gocron_task() {
	fmt.Println("这是gocron_task函数", time.Now().Format("2006-01-02 15:04:05"))
}

func CronDemo() {
	crontab := cron.New(cron.WithSeconds())

	spec := "*/5 * * * * ?" // cron表达式, 每5秒一次

	crontab.AddFunc(spec, gocron_task)

	crontab.Start()

	select {} // 阻塞主线程
}
