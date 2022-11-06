package examples

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
	"testing"
	"time"
)

func newTestLoggerAndPool() (*zap.Logger, *testLogPool) {
	p := newTestLogPool()

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	sink := zapcore.AddSync(p)
	lsink := zapcore.Lock(sink)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), lsink, zapcore.DebugLevel)
	logger := zap.New(core)
	return logger, p
}

func newTestLogPool() *testLogPool {
	return &testLogPool{
		entries: make([]string, 0),
	}
}

type testLogPool struct {
	mu      sync.RWMutex
	entries []string
}

func (w *testLogPool) Write(p []byte) (n int, err error) {
	w.add(strings.TrimRight(string(p), "\n"))
	return len(p), nil
}

func (w *testLogPool) reset() {
	if w.entries == nil {
		return
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	w.entries = w.entries[:0]
}

func (w *testLogPool) last() string {
	if w.entries == nil {
		return ""
	}

	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.entries[len(w.entries)-1]
}

func (w *testLogPool) add(entry string) {
	if w.entries == nil {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	w.entries = append(w.entries, entry)
}

func TestExampleLog_MarshalLogObject(t *testing.T) {
	type fields struct {
		BaseLog       *BaseLog
		BaseArrayLog  BaseArrayLog
		Bool          bool
		BoolPtr       *bool
		Complex128    complex128
		Complex128Ptr *complex128
		Complex64     complex64
		Complex64Ptr  *complex64
		Float64       float64
		Float64Ptr    *float64
		Float32       float32
		Float32Ptr    *float32
		Int           int
		IntPtr        *int
		Int64         int64
		Int64Ptr      *int64
		Int32         int32
		Int32Ptr      *int32
		Int16         int16
		Int16Ptr      *int16
		Int8          int8
		Int8Ptr       *int8
		String        string
		StringPtr     *string
		Uint          uint
		UintPtr       *uint
		Uint64        uint64
		Uint64Ptr     *uint64
		Uint32        uint32
		Uint32Ptr     *uint32
		Uint16        uint16
		Uint16Ptr     *uint16
		Uint8         uint8
		Uint8Ptr      *uint8
		Bytes         []byte
		UintPtrVal    uintptr
		UintPtrValPtr *uintptr
		Time          time.Time
		TimePtr       *time.Time
		Duration      time.Duration
		DurationPtr   *time.Duration
		Map           map[string]bool
		CustomID      string
	}

	baseLog := &BaseLog{Val: 1}
	baseArrayLog := BaseArrayLog{1, 2, 3}

	boolVal := false
	complex128Val := complex128(128)
	complex64Val := complex64(64)
	float64Val := float64(64.64)
	float32Val := float32(32.32)
	intVal := 1
	int64Val := int64(64)
	int32Val := int32(32)
	int16Val := int16(16)
	int8Val := int8(8)
	stringVal := "test"
	uintVal := uint(1)
	uint64Val := uint64(64)
	uint32Val := uint32(32)
	uint16Val := uint16(16)
	uint8Val := uint8(8)
	byteVals := []byte("byte")
	uintptrVal := uintptr(intVal)
	const layout = "2006-01-02 15:04:05"
	timeVal, _ := time.Parse(layout, "2022-11-20 10:01:01")
	durationVal := time.Second
	mapVal := map[string]bool{"key": true}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test_For_ExampleLog",
			fields: fields{
				BaseLog:       baseLog,
				BaseArrayLog:  baseArrayLog,
				Bool:          boolVal,
				BoolPtr:       &boolVal,
				Complex128:    complex128Val,
				Complex128Ptr: &complex128Val,
				Complex64:     complex64Val,
				Complex64Ptr:  &complex64Val,
				Float64:       float64Val,
				Float64Ptr:    &float64Val,
				Float32:       float32Val,
				Float32Ptr:    &float32Val,
				Int:           intVal,
				IntPtr:        &intVal,
				Int64:         int64Val,
				Int64Ptr:      &int64Val,
				Int32:         int32Val,
				Int32Ptr:      &int32Val,
				Int16:         int16Val,
				Int16Ptr:      &int16Val,
				Int8:          int8Val,
				Int8Ptr:       &int8Val,
				String:        stringVal,
				StringPtr:     &stringVal,
				Uint:          uintVal,
				UintPtr:       &uintVal,
				Uint64:        uint64Val,
				Uint64Ptr:     &uint64Val,
				Uint32:        uint32Val,
				Uint32Ptr:     &uint32Val,
				Uint16:        uint16Val,
				Uint16Ptr:     &uint16Val,
				Uint8:         uint8Val,
				Uint8Ptr:      &uint8Val,
				Bytes:         byteVals,
				UintPtrVal:    uintptrVal,
				UintPtrValPtr: &uintptrVal,
				Time:          timeVal,
				TimePtr:       &timeVal,
				Duration:      durationVal,
				DurationPtr:   &durationVal,
				Map:           mapVal,
				CustomID:      stringVal,
			},
			want: "{\"level\":\"info\",\"msg\":\"test\",\"example\":{\"base_log\":{\"Val\":1},\"base_array_log\":[1,2,3],\"bool\":false,\"bool_ptr\":false,\"complex_128\":\"128+0i\",\"complex_128_ptr\":\"128+0i\",\"complex_64\":\"64+0i\",\"complex_64_ptr\":\"64+0i\",\"float_64\":64.64,\"float_64_ptr\":64.64,\"float_32\":32.32,\"float_32_ptr\":32.32,\"int\":1,\"int_ptr\":1,\"int_64\":64,\"int_64_ptr\":64,\"int_32\":32,\"int_32_ptr\":32,\"int_16\":16,\"int_16_ptr\":16,\"int_8\":8,\"int_8_ptr\":8,\"string\":\"test\",\"string_ptr\":\"test\",\"uint\":1,\"uint_ptr\":1,\"uint_64\":64,\"uint_64_ptr\":64,\"uint_32\":32,\"uint_32_ptr\":32,\"uint_16\":16,\"uint_16_ptr\":16,\"uint_8\":8,\"uint_8_ptr\":8,\"bytes\":\"Ynl0ZQ==\",\"uint_ptr_val\":1,\"uint_ptr_val_ptr\":1,\"time\":\"2022-11-20T10:01:01.000Z\",\"time_ptr\":\"2022-11-20T10:01:01.000Z\",\"duration\":\"1s\",\"duration_ptr\":\"1s\",\"map\":{\"key\":true},\"my_id\":\"test\"}}",
		},
	}

	logger, logPool := newTestLoggerAndPool()
	defer logger.Sync()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logPool.reset()
			l := &ExampleLog{
				BaseLog:       tt.fields.BaseLog,
				BaseArrayLog:  tt.fields.BaseArrayLog,
				Bool:          tt.fields.Bool,
				BoolPtr:       tt.fields.BoolPtr,
				Complex128:    tt.fields.Complex128,
				Complex128Ptr: tt.fields.Complex128Ptr,
				Complex64:     tt.fields.Complex64,
				Complex64Ptr:  tt.fields.Complex64Ptr,
				Float64:       tt.fields.Float64,
				Float64Ptr:    tt.fields.Float64Ptr,
				Float32:       tt.fields.Float32,
				Float32Ptr:    tt.fields.Float32Ptr,
				Int:           tt.fields.Int,
				IntPtr:        tt.fields.IntPtr,
				Int64:         tt.fields.Int64,
				Int64Ptr:      tt.fields.Int64Ptr,
				Int32:         tt.fields.Int32,
				Int32Ptr:      tt.fields.Int32Ptr,
				Int16:         tt.fields.Int16,
				Int16Ptr:      tt.fields.Int16Ptr,
				Int8:          tt.fields.Int8,
				Int8Ptr:       tt.fields.Int8Ptr,
				String:        tt.fields.String,
				StringPtr:     tt.fields.StringPtr,
				Uint:          tt.fields.Uint,
				UintPtr:       tt.fields.UintPtr,
				Uint64:        tt.fields.Uint64,
				Uint64Ptr:     tt.fields.Uint64Ptr,
				Uint32:        tt.fields.Uint32,
				Uint32Ptr:     tt.fields.Uint32Ptr,
				Uint16:        tt.fields.Uint16,
				Uint16Ptr:     tt.fields.Uint16Ptr,
				Uint8:         tt.fields.Uint8,
				Uint8Ptr:      tt.fields.Uint8Ptr,
				Bytes:         tt.fields.Bytes,
				UintPtrVal:    tt.fields.UintPtrVal,
				UintPtrValPtr: tt.fields.UintPtrValPtr,
				Time:          tt.fields.Time,
				TimePtr:       tt.fields.TimePtr,
				Duration:      tt.fields.Duration,
				DurationPtr:   tt.fields.DurationPtr,
				Map:           tt.fields.Map,
				CustomID:      tt.fields.CustomID,
			}

			logger.Info("test", zap.Object("example", l))
			got := logPool.last()
			if got != tt.want {
				t.Errorf("log message is wrong. want:%v, got:%v", tt.want, got)
			}
		})
	}
}
