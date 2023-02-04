package logger

import (
	"SiverPineValley/trailer-manager/common"
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
		rotatelogs.WithRotationTime(time.Hour * 24))
	if err != nil {
		panic(err)
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
	encoderConfig.EncodeTime = customLogEncoder
	config.Encoding = "json"
	config.EncoderConfig = encoderConfig

	// initialize the JSON encoding config
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		return
	}

	log, err = config.Build(zap.WrapCore(initRotation), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}