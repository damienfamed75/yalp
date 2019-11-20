package yalp

import "go.uber.org/zap/zapcore"

// interfaceToZapField takes the interfaces passed in and type asserts them
// into a zap.Field and returns a slice.
func interfaceToZapField(iFields ...interface{}) (fields []zapcore.Field) {
	for i := 0; i < len(iFields); i++ {
		fields = append(fields, iFields[i].(zapcore.Field))
	}
	return
}
