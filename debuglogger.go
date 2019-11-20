package yalp

import "go.uber.org/zap"

// DebugLogger is the default logger used for the Quirk client.
// This logger will not print anything out.
type DebugLogger struct {
	logger Logger
}

var _ Logger = &DebugLogger{}

// NewDebugLogger returns a nil logging
// object for the Quirk client to use.
func NewDebugLogger() *DebugLogger {
	l, _ := zap.NewDevelopment()
	return &DebugLogger{
		logger: l,
	}
}

// Info does nothing.
func (l *DebugLogger) Info(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Info(msg, fields...)
}

// Debug logs nothing.
func (l *DebugLogger) Debug(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Debug(msg, fields...)
}

// Error logs nothing.
func (l *DebugLogger) Error(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Error(msg, fields...)
}

// Warn does nothing.
func (l *DebugLogger) Warn(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Warn(msg, fields...)
}

// Fatal logs nothing.
func (l *DebugLogger) Fatal(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Fatal(msg, fields...)
}

// Sugar returns a sugared logger which is typically slower, but nicer to read.
func (l *DebugLogger) Sugar() Logger {
	return &DebugLogger{l.logger.Sugar()}
}
