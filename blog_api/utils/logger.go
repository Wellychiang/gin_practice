package utils

import (
	"api/conf"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var logToFile *logrus.Logger

//日誌文件名
var loggerFile string

func SetLogFile(file string) {
	loggerFile = file
}

//初始化
func init() {
	SetLogFile(filepath.Join(conf.Conf.MyLog.Path, conf.Conf.MyLog.Name))
}

func Log() *logrus.Logger {
	// 文件輸出
	if conf.Conf.MyLog.Model == "file" {
		return logFile()
	} else {
		// 控制台輸出
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2010-01-02 15:02:02"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()

		logToFile.SetLevel(logrus.DebugLevel)

		// 設置 rotatelogs 返回寫日誌對象 logWriter
		logWriter, _ := rotatelogs.New(
			// 分割後的文件名
			loggerFile+"_%Y%m%d.log",

			// 設置最大保存時間
			rotatelogs.WithMaxAge(30*24*time.Hour),

			// 設置日誌切割時間間隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		// 設置時間格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:03:02",
		})

		// 新增 Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile

}
