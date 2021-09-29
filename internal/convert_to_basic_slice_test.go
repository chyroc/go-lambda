package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToIntSlice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []int
		errContain string
	}{
		{name: "ToIntSlice - []int{int(1)}", args: []int{int(1)}, want: []int{int(1)}},
		{name: "ToIntSlice - []interface{}{int(1)}", args: []interface{}{int(1)}, want: []int{int(1)}},
		{name: "ToIntSlice - [2]int{int(1), int(2)}", args: [2]int{int(1), int(2)}, want: []int{int(1), int(2)}},
		{name: "ToIntSlice - [2]interface{}{int(1), int(2)}", args: [2]interface{}{int(1), int(2)}, want: []int{int(1), int(2)}},
		{name: "ToIntSlice - str", args: "str", errContain: "can't convert"},
		{name: "ToIntSlice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntSlice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToInt8Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []int8
		errContain string
	}{
		{name: "ToInt8Slice - []int8{int8(1)}", args: []int8{int8(1)}, want: []int8{int8(1)}},
		{name: "ToInt8Slice - []interface{}{int8(1)}", args: []interface{}{int8(1)}, want: []int8{int8(1)}},
		{name: "ToInt8Slice - [2]int8{int8(1), int8(2)}", args: [2]int8{int8(1), int8(2)}, want: []int8{int8(1), int8(2)}},
		{name: "ToInt8Slice - [2]interface{}{int8(1), int8(2)}", args: [2]interface{}{int8(1), int8(2)}, want: []int8{int8(1), int8(2)}},
		{name: "ToInt8Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToInt8Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToInt16Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []int16
		errContain string
	}{
		{name: "ToInt16Slice - []int16{int16(1)}", args: []int16{int16(1)}, want: []int16{int16(1)}},
		{name: "ToInt16Slice - []interface{}{int16(1)}", args: []interface{}{int16(1)}, want: []int16{int16(1)}},
		{name: "ToInt16Slice - [2]int16{int16(1), int16(1)}", args: [2]int16{int16(1), int16(1)}, want: []int16{int16(1), int16(1)}},
		{name: "ToInt16Slice - [2]interface{}{int16(1), int16(1)}", args: [2]interface{}{int16(1), int16(1)}, want: []int16{int16(1), int16(1)}},
		{name: "ToInt16Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToInt16Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToInt32Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []int32
		errContain string
	}{
		{name: "ToInt32Slice - []int32{int32(1)}", args: []int32{int32(1)}, want: []int32{int32(1)}},
		{name: "ToInt32Slice - []interface{}{int32(1)}", args: []interface{}{int32(1)}, want: []int32{int32(1)}},
		{name: "ToInt32Slice - [2]int32{int32(1), int32(1)}", args: [2]int32{int32(1), int32(1)}, want: []int32{int32(1), int32(1)}},
		{name: "ToInt32Slice - [2]interface{}{int32(1), int32(1)}", args: [2]interface{}{int32(1), int32(1)}, want: []int32{int32(1), int32(1)}},
		{name: "ToInt32Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToInt32Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToInt64Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []int64
		errContain string
	}{
		{name: "ToInt64Slice - []int64{int64(1)}", args: []int64{int64(1)}, want: []int64{int64(1)}},
		{name: "ToInt64Slice - []interface{}{int64(1)}", args: []interface{}{int64(1)}, want: []int64{int64(1)}},
		{name: "ToInt64Slice - [2]int64{int64(1), int64(1)}", args: [2]int64{int64(1), int64(1)}, want: []int64{int64(1), int64(1)}},
		{name: "ToInt64Slice - [2]interface{}{int64(1), int64(1)}", args: [2]interface{}{int64(1), int64(1)}, want: []int64{int64(1), int64(1)}},
		{name: "ToInt64Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToInt64Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUintSlice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uint
		errContain string
	}{
		{name: "ToUintSlice - []uint{uint(1)}", args: []uint{uint(1)}, want: []uint{uint(1)}},
		{name: "ToUintSlice - []interface{}{uint(1)}", args: []interface{}{uint(1)}, want: []uint{uint(1)}},
		{name: "ToUintSlice - [2]uint{uint(1), uint(1)}", args: [2]uint{uint(1), uint(1)}, want: []uint{uint(1), uint(1)}},
		{name: "ToUintSlice - [2]interface{}{uint(1), uint(1)}", args: [2]interface{}{uint(1), uint(1)}, want: []uint{uint(1), uint(1)}},
		{name: "ToUintSlice - str", args: "str", errContain: "can't convert"},
		{name: "ToUintSlice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintSlice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUint8Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uint8
		errContain string
	}{
		{name: "ToUint8Slice - []uint8{uint8(1)}", args: []uint8{uint8(1)}, want: []uint8{uint8(1)}},
		{name: "ToUint8Slice - []interface{}{uint8(1)}", args: []interface{}{uint8(1)}, want: []uint8{uint8(1)}},
		{name: "ToUint8Slice - [2]uint8{uint8(1), uint8(1)}", args: [2]uint8{uint8(1), uint8(1)}, want: []uint8{uint8(1), uint8(1)}},
		{name: "ToUint8Slice - [2]interface{}{uint8(1), uint8(1)}", args: [2]interface{}{uint8(1), uint8(1)}, want: []uint8{uint8(1), uint8(1)}},
		{name: "ToUint8Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToUint8Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUint16Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uint16
		errContain string
	}{
		{name: "ToUint16Slice - []uint16{uint16(1)}", args: []uint16{uint16(1)}, want: []uint16{uint16(1)}},
		{name: "ToUint16Slice - []interface{}{uint16(1)}", args: []interface{}{uint16(1)}, want: []uint16{uint16(1)}},
		{name: "ToUint16Slice - [2]uint16{uint16(1), uint16(1)}", args: [2]uint16{uint16(1), uint16(1)}, want: []uint16{uint16(1), uint16(1)}},
		{name: "ToUint16Slice - [2]interface{}{uint16(1), uint16(1)}", args: [2]interface{}{uint16(1), uint16(1)}, want: []uint16{uint16(1), uint16(1)}},
		{name: "ToUint16Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToUint16Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUint32Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uint32
		errContain string
	}{
		{name: "ToUint32Slice - []uint32{uint32(1)}", args: []uint32{uint32(1)}, want: []uint32{uint32(1)}},
		{name: "ToUint32Slice - []interface{}{uint32(1)}", args: []interface{}{uint32(1)}, want: []uint32{uint32(1)}},
		{name: "ToUint32Slice - [2]uint32{uint32(1), uint32(1)}", args: [2]uint32{uint32(1), uint32(1)}, want: []uint32{uint32(1), uint32(1)}},
		{name: "ToUint32Slice - [2]interface{}{uint32(1), uint32(1)}", args: [2]interface{}{uint32(1), uint32(1)}, want: []uint32{uint32(1), uint32(1)}},
		{name: "ToUint32Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToUint32Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUint64Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uint64
		errContain string
	}{
		{name: "ToUint64Slice - []uint64{uint64(1)}", args: []uint64{uint64(1)}, want: []uint64{uint64(1)}},
		{name: "ToUint64Slice - []interface{}{uint64(1)}", args: []interface{}{uint64(1)}, want: []uint64{uint64(1)}},
		{name: "ToUint64Slice - [2]uint64{uint64(1), uint64(1)}", args: [2]uint64{uint64(1), uint64(1)}, want: []uint64{uint64(1), uint64(1)}},
		{name: "ToUint64Slice - [2]interface{}{uint64(1), uint64(1)}", args: [2]interface{}{uint64(1), uint64(1)}, want: []uint64{uint64(1), uint64(1)}},
		{name: "ToUint64Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToUint64Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToUintptrSlice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []uintptr
		errContain string
	}{
		{name: "ToUintptrSlice - []uintptr{uintptr(1)}", args: []uintptr{uintptr(1)}, want: []uintptr{uintptr(1)}},
		{name: "ToUintptrSlice - []interface{}{uintptr(1)}", args: []interface{}{uintptr(1)}, want: []uintptr{uintptr(1)}},
		{name: "ToUintptrSlice - [2]uintptr{uintptr(1), uintptr(1)}", args: [2]uintptr{uintptr(1), uintptr(1)}, want: []uintptr{uintptr(1), uintptr(1)}},
		{name: "ToUintptrSlice - [2]interface{}{uintptr(1), uintptr(1)}", args: [2]interface{}{uintptr(1), uintptr(1)}, want: []uintptr{uintptr(1), uintptr(1)}},
		{name: "ToUintptrSlice - str", args: "str", errContain: "can't convert"},
		{name: "ToUintptrSlice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintptrSlice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToFloat32Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []float32
		errContain string
	}{
		{name: "ToFloat32Slice - []float32{float32(1)}", args: []float32{float32(1)}, want: []float32{float32(1)}},
		{name: "ToFloat32Slice - []interface{}{float32(1)}", args: []interface{}{float32(1)}, want: []float32{float32(1)}},
		{name: "ToFloat32Slice - [2]float32{float32(1), float32(1)}", args: [2]float32{float32(1), float32(1)}, want: []float32{float32(1), float32(1)}},
		{name: "ToFloat32Slice - [2]interface{}{float32(1), float32(1)}", args: [2]interface{}{float32(1), float32(1)}, want: []float32{float32(1), float32(1)}},
		{name: "ToFloat32Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToFloat32Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToFloat64Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []float64
		errContain string
	}{
		{name: "ToFloat64Slice - []float64{float64(1)}", args: []float64{float64(1)}, want: []float64{float64(1)}},
		{name: "ToFloat64Slice - []interface{}{float64(1)}", args: []interface{}{float64(1)}, want: []float64{float64(1)}},
		{name: "ToFloat64Slice - [2]float64{float64(1), float64(1)}", args: [2]float64{float64(1), float64(1)}, want: []float64{float64(1), float64(1)}},
		{name: "ToFloat64Slice - [2]interface{}{float64(1), float64(1)}", args: [2]interface{}{float64(1), float64(1)}, want: []float64{float64(1), float64(1)}},
		{name: "ToFloat64Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToFloat64Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToBoolSlice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []bool
		errContain string
	}{
		{name: "ToBoolSlice - []bool{bool(true)}", args: []bool{bool(true)}, want: []bool{bool(true)}},
		{name: "ToBoolSlice - []interface{}{bool(true)}", args: []interface{}{bool(true)}, want: []bool{bool(true)}},
		{name: "ToBoolSlice - [2]bool{bool(true), bool(true)}", args: [2]bool{bool(true), bool(true)}, want: []bool{bool(true), bool(true)}},
		{name: "ToBoolSlice - [2]interface{}{bool(true), bool(true)}", args: [2]interface{}{bool(true), bool(true)}, want: []bool{bool(true), bool(true)}},
		{name: "ToBoolSlice - str", args: "str", errContain: "can't convert"},
		{name: "ToBoolSlice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolSlice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToComplex64Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []complex64
		errContain string
	}{
		{name: "ToComplex64Slice - []complex64{complex64(1)}", args: []complex64{complex64(1)}, want: []complex64{complex64(1)}},
		{name: "ToComplex64Slice - []interface{}{complex64(1)}", args: []interface{}{complex64(1)}, want: []complex64{complex64(1)}},
		{name: "ToComplex64Slice - [2]complex64{complex64(1), complex64(1)}", args: [2]complex64{complex64(1), complex64(1)}, want: []complex64{complex64(1), complex64(1)}},
		{name: "ToComplex64Slice - [2]interface{}{complex64(1), complex64(1)}", args: [2]interface{}{complex64(1), complex64(1)}, want: []complex64{complex64(1), complex64(1)}},
		{name: "ToComplex64Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToComplex64Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex64Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToComplex128Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []complex128
		errContain string
	}{
		{name: "ToComplex128Slice - []complex128{complex128(1)}", args: []complex128{complex128(1)}, want: []complex128{complex128(1)}},
		{name: "ToComplex128Slice - []interface{}{complex128(1)}", args: []interface{}{complex128(1)}, want: []complex128{complex128(1)}},
		{name: "ToComplex128Slice - [2]complex128{complex128(1), complex128(1)}", args: [2]complex128{complex128(1), complex128(1)}, want: []complex128{complex128(1), complex128(1)}},
		{name: "ToComplex128Slice - [2]interface{}{complex128(1), complex128(1)}", args: [2]interface{}{complex128(1), complex128(1)}, want: []complex128{complex128(1), complex128(1)}},
		{name: "ToComplex128Slice - str", args: "str", errContain: "can't convert"},
		{name: "ToComplex128Slice - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex128Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}

func Test_ToStringSlice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []string
		errContain string
	}{
		{name: "ToStringSlice - []string{-1-}", args: []string{"1"}, want: []string{"1"}},
		{name: "ToStringSlice - []interface{}{-1-}", args: []interface{}{"1"}, want: []string{"1"}},
		{name: "ToStringSlice - [2]string{-1-, -2-}", args: [2]string{"1", "2"}, want: []string{"1", "2"}},
		{name: "ToStringSlice - [2]interface{}{-1-, -2-}", args: [2]interface{}{"1", "2"}, want: []string{"1", "2"}},
		{name: "ToStringSlice - str", args: "str", errContain: "can't convert"},
		{name: "ToStringSlice - []int{1}", args: []int{1}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringSlice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
