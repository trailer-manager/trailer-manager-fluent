package logger

import (
	"SiverPineValley/trailer-manager/common"
	"SiverPineValley/trailer-manager/utility"
	"context"
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
	if utility.Contains([]string{common.ModeLocal, common.ModeDevelopment, common.ModeStaging}, mode) {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.StacktraceKey = "stacktrace"
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		config.DisableCaller = true
		config.DisableStacktrace = true
	}

	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "msg"
	encoderConfig.CallerKey = "caller"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeTime = customLogEncoder
	//encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	config.Encoding = "json"
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.WrapCore(initRotation), zap.AddCallerSkip(2))
	if err != nil {
		builtin_log.Fatal(err)
	}

	return
}

func makeLogMessage(level string, message string) string {
	lc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(lc).Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf(logFormat, now, level, message)
}

func makeContextLogMessage(level string, ctx context.Context, transactionId string, args ...interface{}) (msg string, fields []zap.Field) {
	fields = make([]zapcore.Field, 0)

	lc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(lc).Format("2006-01-02 15:04:05.000")
	if ctx != nil {
		if logType, ok := ctx.Value(common.ContextLogType).(string); ok && logType != "" {
			fields = append(fields, zap.String("type", logType))
		} else {
			fields = append(fields, zap.String("type", common.ContextLogTypeNormal))
		}
	} else {
		fields = append(fields, zap.String("type", common.ContextLogTypeNormal))
	}
	fields = append(fields, zap.String(common.HeaderTransactionId, transactionId))
	fields = append(fields, zap.String("level", level))
	fields = append(fields, zap.String("time", now))

	msg = fmt.Sprint(args...)
	return msg, fields
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

func InfoContext(ctx context.Context, args ...interface{}) {
	transactionId, _ := ctx.Value(common.HeaderTransactionId).(string)

	msg, fields := makeContextLogMessage("INFO", ctx, transactionId, args)
	log.Info(msg, fields...)
}

func DebugContext(ctx context.Context, args ...interface{}) {
	transactionId, _ := ctx.Value(common.HeaderTransactionId).(string)

	msg, fields := makeContextLogMessage("INFO", ctx, transactionId, args)
	log.Debug(msg, fields...)
}

func ErrorContext(ctx context.Context, args ...interface{}) {
	transactionId, _ := ctx.Value(common.HeaderTransactionId).(string)

	msg, fields := makeContextLogMessage("INFO", ctx, transactionId, args)
	log.Error(msg, fields...)
}

func FatalContext(ctx context.Context, args ...interface{}) {
	transactionId, _ := ctx.Value(common.HeaderTransactionId).(string)

	msg, fields := makeContextLogMessage("INFO", ctx, transactionId, args)
	log.Fatal(msg, fields...)
}
