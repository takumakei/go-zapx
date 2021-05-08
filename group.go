package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// G returns the named group of zap.Field.
func G(name string, field ...zap.Field) zap.Field {
	return zap.Any(name, Fields(field))
}

// Fields is []zap.Field that has the method MarshalLogObject(zapcore.ObjectEncoder) error.
type Fields []zap.Field

func (fs Fields) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for _, v := range fs {
		v.AddTo(enc)
	}
	return nil
}
