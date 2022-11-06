//go:generate go run ../cmd/zmlog/cmd.go -f $GOFILE
package examples

import "time"

type ExampleLog struct {
	*BaseLog
	BaseArrayLog
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
	CustomID      string `key:"my_id"` // custom key
}
