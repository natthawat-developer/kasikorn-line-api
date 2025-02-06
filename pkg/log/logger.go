package log

import (
	"go.uber.org/zap"
)

// Logger is a global logger instance
var Logger *zap.Logger

// Initialize initializes the logger and sets the global Logger variable
func Initialize() {
	var err error
	Logger, err = zap.NewProduction() // หรือใช้ zap.NewDevelopment() สำหรับ development

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
	Logger.Fatal(msg, fields...) // ใช้ Logger ที่เป็น global instance
}

// Error logs a message at Error level
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...) // ใช้ Logger ที่เป็น global instance
}

// Info logs a message at Info level
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...) // ใช้ Logger ที่เป็น global instance
}

// Debug logs a message at Debug level
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...) // ใช้ Logger ที่เป็น global instance
}
