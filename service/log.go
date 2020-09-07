package service

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs" //日志切分
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"runtime"
	"strings"
	"time"
)

type LoggerStruct struct {
	baseLogger *zap.SugaredLogger
}

var Logger, InitLogger *LoggerStruct

var HttpLogger, ErrorLogger *zap.Logger

func InitLog() {

	Logger = &LoggerStruct{
		baseLogger: getLogger("app", "json", zapcore.DebugLevel).Sugar(),
	}
	HttpLogger = getLogger("access", "console", zapcore.InfoLevel)
	ErrorLogger = getLogger("error", "console", zapcore.ErrorLevel)
	InitLogger = &LoggerStruct{
		baseLogger: getLogger("init", "console", zapcore.InfoLevel).Sugar(),
	}
	InitLogger.Info("log", "init log successful.")
}

func getLogger(filename string, encodertype string, level zapcore.Level) *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "keywords",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder, //zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var encoder zapcore.Encoder
	if encodertype == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	logWriter := getWriter("./logs/" + filename)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(logWriter), level),
	)
	return zap.New(core)
}

//日志时间格式化
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1) + "-%Y%m%d.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func LogSync() {
	Logger.baseLogger.Sync()
	HttpLogger.Sync()
	ErrorLogger.Sync()
	InitLogger.baseLogger.Sync()
}

func (this *LoggerStruct) getData(data interface{}) []interface{} {
	slice := make([]interface{}, 0)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file path and line")
	}
	slice = append(slice, "file", fmt.Sprintf("%s:%d", trimmedPath(file), line))
	slice = append(slice, "data", data)
	return slice
}

func trimmedPath(file string) string {

	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return file
	}
	return file[idx+1:]
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Debug(keywords string, value interface{}) {

	this.baseLogger.Debugw(keywords, this.getData(value)...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Info(keywords string, value interface{}) {
	this.baseLogger.Infow(keywords, this.getData(value)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Warn(keywords string, value interface{}) {
	this.baseLogger.Warnw(keywords, this.getData(value)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Error(keywords string, value interface{}) {
	this.baseLogger.Errorw(keywords, this.getData(value)...)
}
