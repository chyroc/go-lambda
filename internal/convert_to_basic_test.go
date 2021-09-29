package internal

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int
		errContain string
	}{
		{name: "ToInt - int(1)", args: int(1), want: int(1)},
		{name: "ToInt - int8(1)", args: int8(1), want: int(1)},
		{name: "ToInt - int16(1)", args: int16(1), want: int(1)},
		{name: "ToInt - int32(1)", args: int32(1), want: int(1)},
		{name: "ToInt - int64(1)", args: int64(1), want: int(1)},
		{name: "ToInt - uint(1)", args: uint(1), want: int(1)},
		{name: "ToInt - uint8(1)", args: uint8(1), want: int(1)},
		{name: "ToInt - uint16(1)", args: uint16(1), want: int(1)},
		{name: "ToInt - uint32(1)", args: uint32(1), want: int(1)},
		{name: "ToInt - uint64(1)", args: uint64(1), want: int(1)},
		{name: "ToInt - uintptr(1)", args: uintptr(1), want: int(1)},
		{name: "ToInt - int(math.MaxInt)", args: int(math.MaxInt), want: int(math.MaxInt)},
		{name: "ToInt - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: int(math.MaxInt8)},
		{name: "ToInt - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: int(math.MaxInt16)},
		{name: "ToInt - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: int(math.MaxInt32)},
		{name: "ToInt - int64(math.MaxInt64)", args: int64(math.MaxInt64), want: int(math.MaxInt64)},
		{name: "ToInt - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToInt - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: int(math.MaxUint8)},
		{name: "ToInt - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: int(math.MaxUint16)},
		{name: "ToInt - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: int(math.MaxUint32)},
		{name: "ToInt - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToInt8(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int8
		errContain string
	}{
		{name: "ToInt8 - int(1)", args: int(1), want: int8(1)},
		{name: "ToInt8 - int8(1)", args: int8(1), want: int8(1)},
		{name: "ToInt8 - int16(1)", args: int16(1), want: int8(1)},
		{name: "ToInt8 - int32(1)", args: int32(1), want: int8(1)},
		{name: "ToInt8 - int64(1)", args: int64(1), want: int8(1)},
		{name: "ToInt8 - uint(1)", args: uint(1), want: int8(1)},
		{name: "ToInt8 - uint8(1)", args: uint8(1), want: int8(1)},
		{name: "ToInt8 - uint16(1)", args: uint16(1), want: int8(1)},
		{name: "ToInt8 - uint32(1)", args: uint32(1), want: int8(1)},
		{name: "ToInt8 - uint64(1)", args: uint64(1), want: int8(1)},
		{name: "ToInt8 - uintptr(1)", args: uintptr(1), want: int8(1)},
		{name: "ToInt8 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToInt8 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: int8(math.MaxInt8)},
		{name: "ToInt8 - int16(math.MaxInt16)", args: int16(math.MaxInt16), errContain: "overflow"},
		{name: "ToInt8 - int32(math.MaxInt32)", args: int32(math.MaxInt32), errContain: "overflow"},
		{name: "ToInt8 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToInt8 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToInt8 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), errContain: "overflow"},
		{name: "ToInt8 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), errContain: "overflow"},
		{name: "ToInt8 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), errContain: "overflow"},
		{name: "ToInt8 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt8 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt8 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToInt16(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int16
		errContain string
	}{
		{name: "ToInt16 - int(1)", args: int(1), want: int16(1)},
		{name: "ToInt16 - int8(1)", args: int8(1), want: int16(1)},
		{name: "ToInt16 - int16(1)", args: int16(1), want: int16(1)},
		{name: "ToInt16 - int32(1)", args: int32(1), want: int16(1)},
		{name: "ToInt16 - int64(1)", args: int64(1), want: int16(1)},
		{name: "ToInt16 - uint(1)", args: uint(1), want: int16(1)},
		{name: "ToInt16 - uint8(1)", args: uint8(1), want: int16(1)},
		{name: "ToInt16 - uint16(1)", args: uint16(1), want: int16(1)},
		{name: "ToInt16 - uint32(1)", args: uint32(1), want: int16(1)},
		{name: "ToInt16 - uint64(1)", args: uint64(1), want: int16(1)},
		{name: "ToInt16 - uintptr(1)", args: uintptr(1), want: int16(1)},
		{name: "ToInt16 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToInt16 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: int16(math.MaxInt8)},
		{name: "ToInt16 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: int16(math.MaxInt16)},
		{name: "ToInt16 - int32(math.MaxInt32)", args: int32(math.MaxInt32), errContain: "overflow"},
		{name: "ToInt16 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToInt16 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToInt16 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: int16(math.MaxUint8)},
		{name: "ToInt16 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), errContain: "overflow"},
		{name: "ToInt16 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), errContain: "overflow"},
		{name: "ToInt16 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt16 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt16 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToInt32(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int32
		errContain string
	}{
		{name: "ToInt32 - int(1)", args: int(1), want: int32(1)},
		{name: "ToInt32 - int8(1)", args: int8(1), want: int32(1)},
		{name: "ToInt32 - int16(1)", args: int16(1), want: int32(1)},
		{name: "ToInt32 - int32(1)", args: int32(1), want: int32(1)},
		{name: "ToInt32 - int64(1)", args: int64(1), want: int32(1)},
		{name: "ToInt32 - uint(1)", args: uint(1), want: int32(1)},
		{name: "ToInt32 - uint8(1)", args: uint8(1), want: int32(1)},
		{name: "ToInt32 - uint16(1)", args: uint16(1), want: int32(1)},
		{name: "ToInt32 - uint32(1)", args: uint32(1), want: int32(1)},
		{name: "ToInt32 - uint64(1)", args: uint64(1), want: int32(1)},
		{name: "ToInt32 - uintptr(1)", args: uintptr(1), want: int32(1)},
		{name: "ToInt32 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToInt32 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: int32(math.MaxInt8)},
		{name: "ToInt32 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: int32(math.MaxInt16)},
		{name: "ToInt32 - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: int32(math.MaxInt32)},
		{name: "ToInt32 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToInt32 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToInt32 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: int32(math.MaxUint8)},
		{name: "ToInt32 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: int32(math.MaxUint16)},
		{name: "ToInt32 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), errContain: "overflow"},
		{name: "ToInt32 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt32 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt32 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToInt64(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int64
		errContain string
	}{
		{name: "ToInt64 - int(1)", args: int(1), want: int64(1)},
		{name: "ToInt64 - int8(1)", args: int8(1), want: int64(1)},
		{name: "ToInt64 - int16(1)", args: int16(1), want: int64(1)},
		{name: "ToInt64 - int32(1)", args: int32(1), want: int64(1)},
		{name: "ToInt64 - int64(1)", args: int64(1), want: int64(1)},
		{name: "ToInt64 - uint(1)", args: uint(1), want: int64(1)},
		{name: "ToInt64 - uint8(1)", args: uint8(1), want: int64(1)},
		{name: "ToInt64 - uint16(1)", args: uint16(1), want: int64(1)},
		{name: "ToInt64 - uint32(1)", args: uint32(1), want: int64(1)},
		{name: "ToInt64 - uint64(1)", args: uint64(1), want: int64(1)},
		{name: "ToInt64 - uintptr(1)", args: uintptr(1), want: int64(1)},
		{name: "ToInt64 - int(math.MaxInt)", args: int(math.MaxInt), want: int64(math.MaxInt)},
		{name: "ToInt64 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: int64(math.MaxInt8)},
		{name: "ToInt64 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: int64(math.MaxInt16)},
		{name: "ToInt64 - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: int64(math.MaxInt32)},
		{name: "ToInt64 - int64(math.MaxInt64)", args: int64(math.MaxInt64), want: int64(math.MaxInt64)},
		{name: "ToInt64 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToInt64 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: int64(math.MaxUint8)},
		{name: "ToInt64 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: int64(math.MaxUint16)},
		{name: "ToInt64 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: int64(math.MaxUint32)},
		{name: "ToInt64 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt64 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToInt64 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUint(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint
		errContain string
	}{
		{name: "ToUint - int(1)", args: int(1), want: uint(1)},
		{name: "ToUint - int8(1)", args: int8(1), want: uint(1)},
		{name: "ToUint - int16(1)", args: int16(1), want: uint(1)},
		{name: "ToUint - int32(1)", args: int32(1), want: uint(1)},
		{name: "ToUint - int64(1)", args: int64(1), want: uint(1)},
		{name: "ToUint - uint(1)", args: uint(1), want: uint(1)},
		{name: "ToUint - uint8(1)", args: uint8(1), want: uint(1)},
		{name: "ToUint - uint16(1)", args: uint16(1), want: uint(1)},
		{name: "ToUint - uint32(1)", args: uint32(1), want: uint(1)},
		{name: "ToUint - uint64(1)", args: uint64(1), want: uint(1)},
		{name: "ToUint - uintptr(1)", args: uintptr(1), want: uint(1)},
		{name: "ToUint - int(math.MaxInt)", args: int(math.MaxInt), want: uint(math.MaxInt)},
		{name: "ToUint - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uint(math.MaxInt8)},
		{name: "ToUint - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: uint(math.MaxInt16)},
		{name: "ToUint - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: uint(math.MaxInt32)},
		{name: "ToUint - uint(math.MaxUint)", args: uint(math.MaxUint), want: uint(math.MaxUint)},
		{name: "ToUint - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uint(math.MaxUint8)},
		{name: "ToUint - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: uint(math.MaxUint16)},
		{name: "ToUint - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: uint(math.MaxUint32)},
		{name: "ToUint - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), want: uint(math.MaxUint64)},
		{name: "ToUint - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), want: uint(math.MaxUint64)},
		{name: "ToUint - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUint8(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint8
		errContain string
	}{
		{name: "ToUint8 - int(1)", args: int(1), want: uint8(1)},
		{name: "ToUint8 - int8(1)", args: int8(1), want: uint8(1)},
		{name: "ToUint8 - int16(1)", args: int16(1), want: uint8(1)},
		{name: "ToUint8 - int32(1)", args: int32(1), want: uint8(1)},
		{name: "ToUint8 - int64(1)", args: int64(1), want: uint8(1)},
		{name: "ToUint8 - uint(1)", args: uint(1), want: uint8(1)},
		{name: "ToUint8 - uint8(1)", args: uint8(1), want: uint8(1)},
		{name: "ToUint8 - uint16(1)", args: uint16(1), want: uint8(1)},
		{name: "ToUint8 - uint32(1)", args: uint32(1), want: uint8(1)},
		{name: "ToUint8 - uint64(1)", args: uint64(1), want: uint8(1)},
		{name: "ToUint8 - uintptr(1)", args: uintptr(1), want: uint8(1)},
		{name: "ToUint8 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToUint8 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uint8(math.MaxInt8)},
		{name: "ToUint8 - int16(math.MaxInt16)", args: int16(math.MaxInt16), errContain: "overflow"},
		{name: "ToUint8 - int32(math.MaxInt32)", args: int32(math.MaxInt32), errContain: "overflow"},
		{name: "ToUint8 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToUint8 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToUint8 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uint8(math.MaxUint8)},
		{name: "ToUint8 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), errContain: "overflow"},
		{name: "ToUint8 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), errContain: "overflow"},
		{name: "ToUint8 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint8 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint8 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUint16(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint16
		errContain string
	}{
		{name: "ToUint16 - int(1)", args: int(1), want: uint16(1)},
		{name: "ToUint16 - int8(1)", args: int8(1), want: uint16(1)},
		{name: "ToUint16 - int16(1)", args: int16(1), want: uint16(1)},
		{name: "ToUint16 - int32(1)", args: int32(1), want: uint16(1)},
		{name: "ToUint16 - int64(1)", args: int64(1), want: uint16(1)},
		{name: "ToUint16 - uint(1)", args: uint(1), want: uint16(1)},
		{name: "ToUint16 - uint8(1)", args: uint8(1), want: uint16(1)},
		{name: "ToUint16 - uint16(1)", args: uint16(1), want: uint16(1)},
		{name: "ToUint16 - uint32(1)", args: uint32(1), want: uint16(1)},
		{name: "ToUint16 - uint64(1)", args: uint64(1), want: uint16(1)},
		{name: "ToUint16 - uintptr(1)", args: uintptr(1), want: uint16(1)},
		{name: "ToUint16 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToUint16 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uint16(math.MaxInt8)},
		{name: "ToUint16 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: uint16(math.MaxInt16)},
		{name: "ToUint16 - int32(math.MaxInt32)", args: int32(math.MaxInt32), errContain: "overflow"},
		{name: "ToUint16 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToUint16 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToUint16 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uint16(math.MaxUint8)},
		{name: "ToUint16 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: uint16(math.MaxUint16)},
		{name: "ToUint16 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), errContain: "overflow"},
		{name: "ToUint16 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint16 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint16 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUint32(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint32
		errContain string
	}{
		{name: "ToUint32 - int(1)", args: int(1), want: uint32(1)},
		{name: "ToUint32 - int8(1)", args: int8(1), want: uint32(1)},
		{name: "ToUint32 - int16(1)", args: int16(1), want: uint32(1)},
		{name: "ToUint32 - int32(1)", args: int32(1), want: uint32(1)},
		{name: "ToUint32 - int64(1)", args: int64(1), want: uint32(1)},
		{name: "ToUint32 - uint(1)", args: uint(1), want: uint32(1)},
		{name: "ToUint32 - uint8(1)", args: uint8(1), want: uint32(1)},
		{name: "ToUint32 - uint16(1)", args: uint16(1), want: uint32(1)},
		{name: "ToUint32 - uint32(1)", args: uint32(1), want: uint32(1)},
		{name: "ToUint32 - uint64(1)", args: uint64(1), want: uint32(1)},
		{name: "ToUint32 - uintptr(1)", args: uintptr(1), want: uint32(1)},
		{name: "ToUint32 - int(math.MaxInt)", args: int(math.MaxInt), errContain: "overflow"},
		{name: "ToUint32 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uint32(math.MaxInt8)},
		{name: "ToUint32 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: uint32(math.MaxInt16)},
		{name: "ToUint32 - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: uint32(math.MaxInt32)},
		{name: "ToUint32 - int64(math.MaxInt64)", args: int64(math.MaxInt64), errContain: "overflow"},
		{name: "ToUint32 - uint(math.MaxUint)", args: uint(math.MaxUint), errContain: "overflow"},
		{name: "ToUint32 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uint32(math.MaxUint8)},
		{name: "ToUint32 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: uint32(math.MaxUint16)},
		{name: "ToUint32 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: uint32(math.MaxUint32)},
		{name: "ToUint32 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint32 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), errContain: "overflow"},
		{name: "ToUint32 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUint64(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint64
		errContain string
	}{
		{name: "ToUint64 - int(1)", args: int(1), want: uint64(1)},
		{name: "ToUint64 - int8(1)", args: int8(1), want: uint64(1)},
		{name: "ToUint64 - int16(1)", args: int16(1), want: uint64(1)},
		{name: "ToUint64 - int32(1)", args: int32(1), want: uint64(1)},
		{name: "ToUint64 - int64(1)", args: int64(1), want: uint64(1)},
		{name: "ToUint64 - uint(1)", args: uint(1), want: uint64(1)},
		{name: "ToUint64 - uint8(1)", args: uint8(1), want: uint64(1)},
		{name: "ToUint64 - uint16(1)", args: uint16(1), want: uint64(1)},
		{name: "ToUint64 - uint32(1)", args: uint32(1), want: uint64(1)},
		{name: "ToUint64 - uint64(1)", args: uint64(1), want: uint64(1)},
		{name: "ToUint64 - uintptr(1)", args: uintptr(1), want: uint64(1)},
		{name: "ToUint64 - int(math.MaxInt)", args: int(math.MaxInt), want: uint64(math.MaxInt)},
		{name: "ToUint64 - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uint64(math.MaxInt8)},
		{name: "ToUint64 - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: uint64(math.MaxInt16)},
		{name: "ToUint64 - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: uint64(math.MaxInt32)},
		{name: "ToUint64 - int64(math.MaxInt64)", args: int64(math.MaxInt64), want: uint64(math.MaxInt64)},
		{name: "ToUint64 - uint(math.MaxUint)", args: uint(math.MaxUint), want: uint64(math.MaxUint)},
		{name: "ToUint64 - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uint64(math.MaxUint8)},
		{name: "ToUint64 - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: uint64(math.MaxUint16)},
		{name: "ToUint64 - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: uint64(math.MaxUint32)},
		{name: "ToUint64 - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), want: uint64(math.MaxUint64)},
		{name: "ToUint64 - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), want: uint64(math.MaxUint64)},
		{name: "ToUint64 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToUintptr(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uintptr
		errContain string
	}{
		{name: "ToUintptr - int(1)", args: int(1), want: uintptr(1)},
		{name: "ToUintptr - int8(1)", args: int8(1), want: uintptr(1)},
		{name: "ToUintptr - int16(1)", args: int16(1), want: uintptr(1)},
		{name: "ToUintptr - int32(1)", args: int32(1), want: uintptr(1)},
		{name: "ToUintptr - int64(1)", args: int64(1), want: uintptr(1)},
		{name: "ToUintptr - uint(1)", args: uint(1), want: uintptr(1)},
		{name: "ToUintptr - uint8(1)", args: uint8(1), want: uintptr(1)},
		{name: "ToUintptr - uint16(1)", args: uint16(1), want: uintptr(1)},
		{name: "ToUintptr - uint32(1)", args: uint32(1), want: uintptr(1)},
		{name: "ToUintptr - uint64(1)", args: uint64(1), want: uintptr(1)},
		{name: "ToUintptr - uintptr(1)", args: uintptr(1), want: uintptr(1)},
		{name: "ToUintptr - int(math.MaxInt)", args: int(math.MaxInt), want: uintptr(math.MaxInt)},
		{name: "ToUintptr - int8(math.MaxInt8)", args: int8(math.MaxInt8), want: uintptr(math.MaxInt8)},
		{name: "ToUintptr - int16(math.MaxInt16)", args: int16(math.MaxInt16), want: uintptr(math.MaxInt16)},
		{name: "ToUintptr - int32(math.MaxInt32)", args: int32(math.MaxInt32), want: uintptr(math.MaxInt32)},
		{name: "ToUintptr - int64(math.MaxInt64)", args: int64(math.MaxInt64), want: uintptr(math.MaxInt64)},
		{name: "ToUintptr - uint(math.MaxUint)", args: uint(math.MaxUint), want: uintptr(math.MaxUint)},
		{name: "ToUintptr - uint8(math.MaxUint8)", args: uint8(math.MaxUint8), want: uintptr(math.MaxUint8)},
		{name: "ToUintptr - uint16(math.MaxUint16)", args: uint16(math.MaxUint16), want: uintptr(math.MaxUint16)},
		{name: "ToUintptr - uint32(math.MaxUint32)", args: uint32(math.MaxUint32), want: uintptr(math.MaxUint32)},
		{name: "ToUintptr - uint64(math.MaxUint64)", args: uint64(math.MaxUint64), want: uintptr(math.MaxUint64)},
		{name: "ToUintptr - uintptr(math.MaxUint64)", args: uintptr(math.MaxUint64), want: uintptr(math.MaxUint64)},
		{name: "ToUintptr - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintptr(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToFloat32(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       float32
		errContain string
	}{
		{name: "ToFloat32 - float32(1)", args: float32(1), want: float32(1)},
		{name: "ToFloat32 - float64(1)", args: float64(1), want: float32(1)},
		{name: "ToFloat32 - float32(math.MaxFloat32)", args: float32(math.MaxFloat32), want: float32(math.MaxFloat32)},
		{name: "ToFloat32 - float64(math.MaxFloat64)", args: float64(math.MaxFloat64), errContain: "overflow"},
		{name: "ToFloat32 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToFloat64(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       float64
		errContain string
	}{
		{name: "ToFloat64 - float32(1)", args: float32(1), want: float64(1)},
		{name: "ToFloat64 - float64(1)", args: float64(1), want: float64(1)},
		{name: "ToFloat64 - float32(math.MaxFloat32)", args: float32(math.MaxFloat32), want: float64(math.MaxFloat32)},
		{name: "ToFloat64 - float64(math.MaxFloat64)", args: float64(math.MaxFloat64), want: float64(math.MaxFloat64)},
		{name: "ToFloat64 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToBool(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       bool
		errContain string
	}{
		{name: "ToBool - true", args: true, want: true},
		{name: "ToBool - false", args: false, want: false},
		{name: "ToBool - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToComplex64(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       complex64
		errContain string
	}{
		{name: "ToComplex64 - complex64(1)", args: complex64(1), want: complex64(1)},
		{name: "ToComplex64 - complex128(1)", args: complex128(1), errContain: "can't convert"},
		{name: "ToComplex64 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex64(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToComplex128(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       complex128
		errContain string
	}{
		{name: "ToComplex128 - complex64(1)", args: complex64(1), want: complex128(1)},
		{name: "ToComplex128 - complex128(1)", args: complex128(1), want: complex128(1)},
		{name: "ToComplex128 - str", args: "str", errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex128(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func TestToString(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       string
		errContain string
	}{
		{name: "ToString - string(1)", args: string("1"), want: string("1")},
		{name: "ToString - []rune(1)", args: []rune("1"), want: string("1")},
		{name: "ToString - []byte(1)", args: []byte("1"), want: string("1")},
		{name: "ToString - 1", args: 1, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
