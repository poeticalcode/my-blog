package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 此处用到另外一个第三方库lumberjack，可以实现日志切割等功能，具体的配置项可自行配置，此处是输出json格式的日志
func InitZap() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "request-time",
		LevelKey:       "level",
		NameKey:        "Logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename: "test.log",
		MaxSize:  10,
	}
	writeSyncer := zapcore.AddSync(lumberjackLogger)
	var l zapcore.Level
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start) // 本次请求的总共消耗时间
		// 写入日志
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
