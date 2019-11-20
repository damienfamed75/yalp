package yalp

import "go.uber.org/zap"

// SugarLogger is the default logger used for the Quirk client.
// This logger will not print anything out.
type SugarLogger struct {
	logger *zap.SugaredLogger
}

var _ Logger = &SugarLogger{}

// NewSugarLogger returns a nil logging
// object for the Quirk client to use.
func NewSugarLogger() *SugarLogger {
	l, _ := zap.NewDevelopment()
	ll := l.Sugar()
	return &SugarLogger{
		logger: ll,
	}
}

// Info does nothing.
func (l *SugarLogger) Info(msg string, iFields ...interface{}) {
	l.logger.Infof(msg, iFields...)
}

// Debug logs nothing.
func (l *SugarLogger) Debug(msg string, iFields ...interface{}) {
	l.logger.Debugf(msg, iFields...)
}

// Error logs nothing.
func (l *SugarLogger) Error(msg string, iFields ...interface{}) {
	l.logger.Errorf(msg, iFields...)
}

// Warn does nothing.
func (l *SugarLogger) Warn(msg string, iFields ...interface{}) {
	l.logger.Warnf(msg, iFields...)
}

// Fatal logs nothing.
func (l *SugarLogger) Fatal(msg string, iFields ...interface{}) {
	l.logger.Fatalf(msg, iFields...)
}
