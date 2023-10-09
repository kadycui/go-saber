package logger

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(filename string) {
	// 创建一个新的logrus实例
	Log = logrus.New()

	// 设置日志输出为指定的文件
	file := &lumberjack.Logger{
		Filename:   getLogFilename(filename),
		MaxSize:    100,  // 单位为MB，超出大小则自动分割
		MaxBackups: 3,    // 分割后最多保留的文件数
		MaxAge:     30,   // 保留的最大天数
		LocalTime:  true, // 使用本地时间作为日志文件名和切割时间
	}

	// 打开日志文件
	//file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	Log.Fatal(err)
	//}
	Log.SetOutput(file)

	// 设置日志格式为JSON格式
	Log.SetFormatter(&logrus.JSONFormatter{})
}

func Info(msg string) {
	Log.Info(msg)
}

func Error(msg string) {
	Log.Error(msg)
}

func getLogFilename(filename string) string {
	t := time.Now()
	fmt.Println(filename)
	filename = filename + t.Format("2006-01-02") + ".log"
	fmt.Println(filename)
	return filename
}
