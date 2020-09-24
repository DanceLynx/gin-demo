package log

import (
	"go.uber.org/zap"
)

var HttpLogger, ErrorLogger, InitLogger, GormLogger *zap.Logger

func initInternalLog() {
	GormLogger = getLogger("mysql", "json", "debug")
	HttpLogger = getLogger("request", "console", "info")
	ErrorLogger = getLogger("error", "console", "error")
	InitLogger = getLogger("init", "console", "info")
}

func flushInternalLog() {
	GormLogger.Sync()
	HttpLogger.Sync()
	ErrorLogger.Sync()
	InitLogger.Sync()
}
