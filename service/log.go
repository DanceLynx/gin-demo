package service

import (
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs" //日志切分
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hello/config"
	"io"
	"runtime"
	"strings"
	"time"
)

type LoggerStruct struct {
	baseLogger *zap.SugaredLogger
}

func getLogLevel(level string) zapcore.Level {

	m := map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"error": zapcore.ErrorLevel,
		"warn":  zapcore.WarnLevel,
	}

	l, ok := m[level]
	if ok {
		return l
	} else {
		return zapcore.InfoLevel
	}
}

var Logger, gormLogger *LoggerStruct

var HttpLogger, ErrorLogger, InitLogger *zap.Logger

func InitLog() {

	Logger = &LoggerStruct{
		baseLogger: getLogger("app", config.Log.Formatter, config.Log.LogLevel).Sugar(),
	}
	gormLogger = &LoggerStruct{
		baseLogger: getLogger("mysql", "json", "debug").Sugar(),
	}
	HttpLogger = getLogger("request", "console", "info")
	ErrorLogger = getLogger("error", "console", "error")
	InitLogger = getLogger("init", "console", "info")
	InitLogger.Info("init log successful.")
}

func getLogger(filename string, encodertype string, level string) *zap.Logger {

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
	if encodertype == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	logWriter := getWriter("./logs/" + filename)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(logWriter), getLogLevel(level)),
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
	gormLogger.baseLogger.Sync()
	HttpLogger.Sync()
	ErrorLogger.Sync()
	InitLogger.Sync()
}

func (this *LoggerStruct) getData(ctx context.Context, data interface{}) []interface{} {
	slice := make([]interface{}, 0)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file path and line")
	}

	fileAndLine := fmt.Sprintf("%s:%d", file, line)
	//gorm的file写在data里面，替换到外面标准格式
	if m, ok := data.(map[string]interface{}); ok {
		value, ok := m["file"]
		if ok {
			fileAndLine = value.(string)
			delete(m, "file")

		}
	}
	if ctx.Value("traceId") != nil {
		slice = append(slice, "traceId", ctx.Value("traceId"))
		slice = append(slice, "file", trimmedPath(fileAndLine))
	}
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
func (this *LoggerStruct) Debug(ctx context.Context, keywords string, value interface{}) {

	this.baseLogger.Debugw(keywords, this.getData(ctx, value)...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Info(ctx context.Context, keywords string, value interface{}) {
	this.baseLogger.Infow(keywords, this.getData(ctx, value)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Warn(ctx context.Context, keywords string, value interface{}) {
	this.baseLogger.Warnw(keywords, this.getData(ctx, value)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (this *LoggerStruct) Error(ctx context.Context, keywords string, value interface{}) {
	this.baseLogger.Errorw(keywords, this.getData(ctx, value)...)
}
