package log

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)


var Logger *zap.Logger


type ZapGormLogger struct {
	*zap.Logger
}

// Initialize initializes the logger and sets the global Logger variable
func Initialize() {
	var err error
	Logger, err = zap.NewProduction() 

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

// Close closes the logger
func Close() {
	if err := Logger.Sync(); err != nil {
		panic("Failed to sync logger: " + err.Error())
	}
}

// Fatal logs a message at Fatal level and stops the application
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...) 
}

// Error logs a message at Error level
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Info logs a message at Info level
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Debug logs a message at Debug level
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...) 
}


func NewZapGormLogger(logger *zap.Logger) *ZapGormLogger {
	return &ZapGormLogger{Logger: logger}
}


func (z *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return z
}

// Info logs a message at Info level for GORM
func (z *ZapGormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	z.Logger.Sugar().Infof(msg, args...)
}

// Warn logs a message at Warn level for GORM
func (z *ZapGormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	z.Logger.Sugar().Warnf(msg, args...)
}

// Error logs a message at Error level for GORM
func (z *ZapGormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	z.Logger.Sugar().Errorf(msg, args...)
}

// Trace logs a message at Trace level for GORM
func (z *ZapGormLogger) Trace(ctx context.Context, start time.Time, fc func() (string, int64), err error) {
	duration := time.Since(start)
	sql, rows := fc()

	if err != nil {
		z.Logger.Sugar().Errorf("[TRACE] Error: %s, SQL: %s, Rows: %d, Duration: %s", err.Error(), sql, rows, duration)
	} else {
		z.Logger.Sugar().Debugf("[TRACE] SQL: %s, Rows: %d, Duration: %s", sql, rows, duration)
	}
}
