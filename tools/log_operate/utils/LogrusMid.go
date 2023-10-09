package utils

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type MyLogger struct {
	logger *logrus.Entry
}

func New() MyLogger {

	path := "/home/debian/code/gocode/github.com/kadycui/go-saber/log_operate/go.log"

	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/

	logger := logrus.New()
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		// rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logger.SetOutput(writer)

	// logger.SetLevel(logrus.InfoLevel)
	// logger.SetReportCaller(true)

	// logger.SetOutput(os.Stdout)

	// logger.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02T15:04:05",
	// })

	return MyLogger{
		logger: logger.WithFields(logrus.Fields{"app": "test_demo"}),
	}
}

func (m MyLogger) Trace(args ...interface{}) {
	m.logger.Trace(args...)
}

func (m MyLogger) Tracef(format string, args ...interface{}) {
	m.logger.Tracef(format, args...)
}

func (m MyLogger) Info(args ...interface{}) {
	m.logger.Info(args...)
}

func (m MyLogger) Infof(format string, args ...interface{}) {
	m.logger.Infof(format, args...)
}

func (m MyLogger) Debug(args ...interface{}) {
	m.logger.Debug(args...)
}

func (m MyLogger) Debugf(format string, args ...interface{}) {
	m.logger.Debugf(format, args...)
}

func (m MyLogger) Warning(args ...interface{}) {
	m.logger.Warning(args...)
}

func (m MyLogger) Warningf(format string, args ...interface{}) {
	m.logger.Warningf(format, args...)
}

func (m MyLogger) Error(args ...interface{}) {
	m.logger.Error(args...)
}
