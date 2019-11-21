package yalp

import (
	"go.uber.org/zap"
)

// ConfiguredLogger is the default logger used for the Quirk client.
// This logger will not print anything out.
type ConfiguredLogger struct {
	logger *zap.Logger
}

var _ Logger = &ConfiguredLogger{}

// NewConfiguredLogger returns a nil logging
// object for the Quirk client to use.
func NewConfiguredLogger(cfg zap.Config) *ConfiguredLogger {
	l, _ := cfg.Build()
	return &ConfiguredLogger{
		logger: l,
	}
}

// Info does nothing.
func (l *ConfiguredLogger) Info(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Info(msg, fields...)
}

// Debug logs nothing.
func (l *ConfiguredLogger) Debug(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Debug(msg, fields...)
}

// Error logs nothing.
func (l *ConfiguredLogger) Error(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Error(msg, fields...)
}

// Warn does nothing.
func (l *ConfiguredLogger) Warn(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Warn(msg, fields...)
}

// Fatal logs nothing.
func (l *ConfiguredLogger) Fatal(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Fatal(msg, fields...)
}

func (l *ConfiguredLogger) Sync() error {
	return l.logger.Sync()
}
