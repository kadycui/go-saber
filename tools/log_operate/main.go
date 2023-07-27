package main

import (
	"gostart/log_operate/utils"
	"time"
)

func main() {
	logger := utils.New()
	for {
		logger.Info("11111111111")
		logger.Debug("2222222222")
		logger.Warning("333333333")
		logger.Error("444444444444")
		time.Sleep(time.Duration(10) * time.Second)
	}

}
