package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
	"time"
)

const (
	ERROR_OUTPUT_PATH = "logs/error.log"
	LOG_OUTPUT_PATH	  = "logs/ms_go.log"
	ENCODING 		  = "console"
	TIMESTAMP_FORMAT  = "2006-01-02T15:04:05"
)

var Logger *zap.Logger

func InitLogger(logLevel string) {
    var config zap.Config
	config = zap.NewProductionConfig()
	setOutputPaths(&config)
    setLoggerLevel(logLevel, &config)
	setEnconderConfig(&config)
	buildConfiguration(&config)
}

func setOutputPaths(config *zap.Config) {
	config.OutputPaths = []string{LOG_OUTPUT_PATH}
	config.ErrorOutputPaths = []string{ERROR_OUTPUT_PATH} 
}

func setLoggerLevel(logLevel string, config *zap.Config) {
	// Configuração do nível de log
    switch logLevel {
    case "debug":
        config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    case "info":
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    case "warn":
        config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
    case "error":
        config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
    case "fatal":
        config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
    default:
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    }
}

func setEnconderConfig(config *zap.Config) {
	config.Encoding = ENCODING
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    config.EncoderConfig.EncodeTime = customTimeEncoder
}

func buildConfiguration(config *zap.Config) {
	var err error
    Logger, err = config.Build()
    if err != nil {
        panic("failed to initialize logger")
    }
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format(TIMESTAMP_FORMAT + " "))
}

func Info(msg string, fields ...interface{}) {
    Logger.Info(msg, convertToZapFields(fields)...)
}

func Debug(msg string, fields ...interface{}) {
    Logger.Debug(msg, convertToZapFields(fields)...)
}

func Warn(msg string, fields ...interface{}) {
    Logger.Warn(msg, convertToZapFields(fields)...)
}

func Error(msg string, fields ...interface{}) {
    Logger.Error(msg, convertToZapFields(fields)...)
}

func Fatal(msg string, fields ...interface{}) {
    Logger.Fatal(msg, convertToZapFields(fields)...)
}

func convertToZapFields(fields []interface{}) []zap.Field {
    zapFields := make([]zap.Field, len(fields))
    for i, field := range fields {
        switch f := field.(type) {
        case zap.Field:
            zapFields[i] = f
        case error:
            zapFields[i] = zap.Error(f)
        case string:
            zapFields[i] = zap.String("message", f)
        case int:
            zapFields[i] = zap.Int("interger", f)
        case float64:
            zapFields[i] = zap.Float64("decimal", f)
        case bool:
            zapFields[i] = zap.Bool("boolean", f)
        default:
            zapFields[i] = zap.Any("custom", f)
        }
    }
    return zapFields
}