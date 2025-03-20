package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.SugaredLogger
)

// code from here https://github.com/nlypage/webTemplate/blob/main/internal/adapters/logger/logger.go

// New is a function to initialize logger
/*
 * debug bool - is debug mode
 */
func New(debug bool) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "timestamp",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalColorLevelEncoder, // Цветная подсветка уровней
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(time.Local).Format("2006-01-02 15:04:05"))
		}, // Кастомный формат времени
		EncodeCaller:   zapcore.ShortCallerEncoder, // Краткий формат caller
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	var level zapcore.Level
	if debug {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level)
	log := zap.New(core, zap.AddCaller())

	Log = log.Sugar()
}
