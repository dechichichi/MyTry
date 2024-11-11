package log

import (
	log "github.com/sirupsen/logrus"
)

func Logger() {
	logger := &lumberjack.Logger{
		//路径
		Filename: "gin.log",
		//最大size
		MaxSize: 100, // megabytes
		//最大过期日志保留个数
		MaxBackups: 3,
		//保留过期日志的天数
		MaxAge: 30, //days
		//是否压缩
		Compress: true,
	}
	log.SetOutput(logger)

}
