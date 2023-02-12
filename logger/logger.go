package logger

import (
	"SiverPineValley/trailer-manager/common"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	builtin_log "log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	logFormat        = "%s %s --- %s"
	logFormatContext = "%s %s %s [%s, %s] --- %s"
)

var log *zap.Logger

func customLogEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	sec := float64(nanos) / float64(time.Second)
	enc.AppendFloat64(sec)
}

func initRotation(c zapcore.Core) zapcore.Core {
	// initialize the rotator
	logFile := "/var/log/tm/trailer-manager-%Y-%m-%d.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		builtin_log.Fatal(err)
	}

	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	cores := zapcore.NewTee(c, core)

	return cores
}

func InitLogger(mode string) (err error) {
	config := zap.NewProductionConfig()

	encoderConfig := zapcore.EncoderConfig{}
	if mode == common.ModeDevelopment || mode == common.ModeStaging {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.StacktraceKey = "stacktrace"
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		config.DisableCaller = true
		config.DisableStacktrace = true
	}

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "caller"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeTime = customLogEncoder
	//encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	config.Encoding = "json"
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.WrapCore(initRotation), zap.AddCallerSkip(1))
	if err != nil {
		builtin_log.Fatal(err)
	}

	return
}

func makeLogMessage(level string, message string) string {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf(logFormat, now, level, message)
}

func makeContextLogMessage(level string, ctx interface{}, traceId, source, target, message string) string {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf(logFormatContext, now, level, traceId, source, target, message)
}

func Info(message string, fields ...zap.Field) {
	msg := makeLogMessage("INFO", message)
	log.Info(msg, fields...)
}

func Debug(message string, fields ...zap.Field) {
	msg := makeLogMessage("DEBUG", message)
	log.Debug(msg, fields...)
}

func Error(message string, fields ...zap.Field) {
	msg := makeLogMessage("ERROR", message)
	log.Error(msg, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	msg := makeLogMessage("FATAL", message)
	log.Fatal(msg, fields...)
}

func InfoContext(ctx interface{}, traceId, source, target, message string, fields ...zap.Field) {
	msg := makeContextLogMessage("INFO", ctx, traceId, source, target, message)
	log.Info(msg, fields...)
}

func DebugContext(ctx interface{}, traceId, source, target, message string, fields ...zap.Field) {
	msg := makeContextLogMessage("INFO", ctx, traceId, source, target, message)
	log.Debug(msg, fields...)
}

func ErrorContext(ctx interface{}, traceId, source, target, message string, fields ...zap.Field) {
	msg := makeContextLogMessage("INFO", ctx, traceId, source, target, message)
	log.Error(msg, fields...)
}

func FatalContext(ctx interface{}, traceId, source, target, message string, fields ...zap.Field) {
	msg := makeContextLogMessage("INFO", ctx, traceId, source, target, message)
	log.Fatal(msg, fields...)
}
