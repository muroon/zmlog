package examples

import "go.uber.org/zap/zapcore"

// BaseLog the struct to be embedded in ExampleLog
type BaseLog struct {
	Val int
}

func (l *BaseLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("val", l.Val)
	return nil
}

// BaseArrayLog the struct to be embedded in ExampleLog
type BaseArrayLog []int

func (l BaseArrayLog) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, v := range l {
		enc.AppendInt(v)
	}
	return nil
}
