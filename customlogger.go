package yalp

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CustomLogger is a customizable logging.Logger where you can choose the level
// and the zapcore encoder configuration.
type CustomLogger struct {
	logger Logger
	level  zap.AtomicLevel
	output *os.File
	config zapcore.EncoderConfig
}

var _ Logger = &CustomLogger{}

// NewCustomLogger returns a custom logging
// object for the Classy service to use.
func NewCustomLogger(level []byte, config zapcore.EncoderConfig) *CustomLogger {
	logLevel := zap.NewAtomicLevel()
	logLevel.UnmarshalText(level)

	return &CustomLogger{
		level:  logLevel,
		config: config,
		output: os.Stdout,
		logger: zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.Lock(os.Stdout), logLevel)),
	}
}

// Info logs at an info level.
func (l *CustomLogger) Info(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Info(msg, fields...)
}

// Debug logs at an debug level.
func (l *CustomLogger) Debug(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Debug(msg, fields...)
}

// Warn warns the client.
func (l *CustomLogger) Warn(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Warn(msg, fields...)
}

// Error logs at an error level.
func (l *CustomLogger) Error(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Error(msg, fields...)
}

// Fatal logs at a fatal level and exits.
func (l *CustomLogger) Fatal(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.(*zap.Logger).Fatal(msg, fields...)
}

// SetLevel changes the logger level
func (l *CustomLogger) SetLevel(level []byte) {
	// flush the existing logger before changing to new log level
	l.logger.(*zap.Logger).Sync()

	// Read in the new zapcore AtomicLevel and apply new zap instance
	l.level.UnmarshalText(level)
	l.logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(l.config), zapcore.Lock(l.output), l.level))
}

// SetOutput changes the output
func (l *CustomLogger) SetOutput(output *os.File) {
	// flush the existing logger before changing to new log output
	l.logger.(*zap.Logger).Sync()

	// set the new output and apply new zap instance
	l.output = output
	l.logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(l.config), zapcore.Lock(l.output), l.level))
}

func (l *CustomLogger) Sugar() Logger {
	return &CustomLogger{
		logger: l.logger.Sugar(),
		level:  l.level,
		output: l.output,
		config: l.config,
	}
}
