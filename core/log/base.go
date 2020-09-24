package log

import (
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs" //日志切分
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

func InitLog() {
	initAppLog()
	initInternalLog()
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
	flushAppLog()
	flushInternalLog()
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
