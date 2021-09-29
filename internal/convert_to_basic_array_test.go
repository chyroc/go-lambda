package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToIntArray(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToIntArray - []int{int(1)}", args: []int{int(1)}, want: [1]int{int(1)}},
		{name: "ToIntArray - []interface{}{int(1)}", args: []interface{}{int(1)}, want: [1]int{int(1)}},
		{name: "ToIntArray - [2]int{int(1), int(2)}", args: [2]int{int(1), int(2)}, want: [2]int{int(1), int(2)}},
		{name: "ToIntArray - [2]interface{}{int(1), int(2)}", args: [2]interface{}{int(1), int(2)}, want: [2]int{int(1), int(2)}},
		{name: "ToIntArray - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToIntArray - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntArray(tt.args)
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

func Test_ToInt8Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToInt8Array - []int8{int8(1)}", args: []int8{int8(1)}, want: [1]int8{int8(1)}},
		{name: "ToInt8Array - []interface{}{int8(1)}", args: []interface{}{int8(1)}, want: [1]int8{int8(1)}},
		{name: "ToInt8Array - [2]int8{int8(1), int8(2)}", args: [2]int8{int8(1), int8(2)}, want: [2]int8{int8(1), int8(2)}},
		{name: "ToInt8Array - [2]interface{}{int8(1), int8(2)}", args: [2]interface{}{int8(1), int8(2)}, want: [2]int8{int8(1), int8(2)}},
		{name: "ToInt8Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToInt8Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8Array(tt.args)
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

func Test_ToInt16Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToInt16Array - []int16{int16(1)}", args: []int16{int16(1)}, want: [1]int16{int16(1)}},
		{name: "ToInt16Array - []interface{}{int16(1)}", args: []interface{}{int16(1)}, want: [1]int16{int16(1)}},
		{name: "ToInt16Array - [2]int16{int16(1), int16(1)}", args: [2]int16{int16(1), int16(1)}, want: [2]int16{int16(1), int16(1)}},
		{name: "ToInt16Array - [2]interface{}{int16(1), int16(1)}", args: [2]interface{}{int16(1), int16(1)}, want: [2]int16{int16(1), int16(1)}},
		{name: "ToInt16Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToInt16Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16Array(tt.args)
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

func Test_ToInt32Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToInt32Array - []int32{int32(1)}", args: []int32{int32(1)}, want: [1]int32{int32(1)}},
		{name: "ToInt32Array - []interface{}{int32(1)}", args: []interface{}{int32(1)}, want: [1]int32{int32(1)}},
		{name: "ToInt32Array - [2]int32{int32(1), int32(1)}", args: [2]int32{int32(1), int32(1)}, want: [2]int32{int32(1), int32(1)}},
		{name: "ToInt32Array - [2]interface{}{int32(1), int32(1)}", args: [2]interface{}{int32(1), int32(1)}, want: [2]int32{int32(1), int32(1)}},
		{name: "ToInt32Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToInt32Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32Array(tt.args)
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

func Test_ToInt64Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToInt64Array - []int64{int64(1)}", args: []int64{int64(1)}, want: [1]int64{int64(1)}},
		{name: "ToInt64Array - []interface{}{int64(1)}", args: []interface{}{int64(1)}, want: [1]int64{int64(1)}},
		{name: "ToInt64Array - [2]int64{int64(1), int64(1)}", args: [2]int64{int64(1), int64(1)}, want: [2]int64{int64(1), int64(1)}},
		{name: "ToInt64Array - [2]interface{}{int64(1), int64(1)}", args: [2]interface{}{int64(1), int64(1)}, want: [2]int64{int64(1), int64(1)}},
		{name: "ToInt64Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToInt64Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64Array(tt.args)
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

func Test_ToUintArray(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUintArray - []uint{uint(1)}", args: []uint{uint(1)}, want: [1]uint{uint(1)}},
		{name: "ToUintArray - []interface{}{uint(1)}", args: []interface{}{uint(1)}, want: [1]uint{uint(1)}},
		{name: "ToUintArray - [2]uint{uint(1), uint(1)}", args: [2]uint{uint(1), uint(1)}, want: [2]uint{uint(1), uint(1)}},
		{name: "ToUintArray - [2]interface{}{uint(1), uint(1)}", args: [2]interface{}{uint(1), uint(1)}, want: [2]uint{uint(1), uint(1)}},
		{name: "ToUintArray - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUintArray - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintArray(tt.args)
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

func Test_ToUint8Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUint8Array - []uint8{uint8(1)}", args: []uint8{uint8(1)}, want: [1]uint8{uint8(1)}},
		{name: "ToUint8Array - []interface{}{uint8(1)}", args: []interface{}{uint8(1)}, want: [1]uint8{uint8(1)}},
		{name: "ToUint8Array - [2]uint8{uint8(1), uint8(1)}", args: [2]uint8{uint8(1), uint8(1)}, want: [2]uint8{uint8(1), uint8(1)}},
		{name: "ToUint8Array - [2]interface{}{uint8(1), uint8(1)}", args: [2]interface{}{uint8(1), uint8(1)}, want: [2]uint8{uint8(1), uint8(1)}},
		{name: "ToUint8Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUint8Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8Array(tt.args)
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

func Test_ToUint16Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUint16Array - []uint16{uint16(1)}", args: []uint16{uint16(1)}, want: [1]uint16{uint16(1)}},
		{name: "ToUint16Array - []interface{}{uint16(1)}", args: []interface{}{uint16(1)}, want: [1]uint16{uint16(1)}},
		{name: "ToUint16Array - [2]uint16{uint16(1), uint16(1)}", args: [2]uint16{uint16(1), uint16(1)}, want: [2]uint16{uint16(1), uint16(1)}},
		{name: "ToUint16Array - [2]interface{}{uint16(1), uint16(1)}", args: [2]interface{}{uint16(1), uint16(1)}, want: [2]uint16{uint16(1), uint16(1)}},
		{name: "ToUint16Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUint16Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16Array(tt.args)
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

func Test_ToUint32Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUint32Array - []uint32{uint32(1)}", args: []uint32{uint32(1)}, want: [1]uint32{uint32(1)}},
		{name: "ToUint32Array - []interface{}{uint32(1)}", args: []interface{}{uint32(1)}, want: [1]uint32{uint32(1)}},
		{name: "ToUint32Array - [2]uint32{uint32(1), uint32(1)}", args: [2]uint32{uint32(1), uint32(1)}, want: [2]uint32{uint32(1), uint32(1)}},
		{name: "ToUint32Array - [2]interface{}{uint32(1), uint32(1)}", args: [2]interface{}{uint32(1), uint32(1)}, want: [2]uint32{uint32(1), uint32(1)}},
		{name: "ToUint32Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUint32Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32Array(tt.args)
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

func Test_ToUint64Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUint64Array - []uint64{uint64(1)}", args: []uint64{uint64(1)}, want: [1]uint64{uint64(1)}},
		{name: "ToUint64Array - []interface{}{uint64(1)}", args: []interface{}{uint64(1)}, want: [1]uint64{uint64(1)}},
		{name: "ToUint64Array - [2]uint64{uint64(1), uint64(1)}", args: [2]uint64{uint64(1), uint64(1)}, want: [2]uint64{uint64(1), uint64(1)}},
		{name: "ToUint64Array - [2]interface{}{uint64(1), uint64(1)}", args: [2]interface{}{uint64(1), uint64(1)}, want: [2]uint64{uint64(1), uint64(1)}},
		{name: "ToUint64Array - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUint64Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64Array(tt.args)
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

func Test_ToUintptrArray(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToUintptrArray - []uintptr{uintptr(1)}", args: []uintptr{uintptr(1)}, want: [1]uintptr{uintptr(1)}},
		{name: "ToUintptrArray - []interface{}{uintptr(1)}", args: []interface{}{uintptr(1)}, want: [1]uintptr{uintptr(1)}},
		{name: "ToUintptrArray - [2]uintptr{uintptr(1), uintptr(1)}", args: [2]uintptr{uintptr(1), uintptr(1)}, want: [2]uintptr{uintptr(1), uintptr(1)}},
		{name: "ToUintptrArray - [2]interface{}{uintptr(1), uintptr(1)}", args: [2]interface{}{uintptr(1), uintptr(1)}, want: [2]uintptr{uintptr(1), uintptr(1)}},
		{name: "ToUintptrArray - int(1)", args: int(1), errContain: "can't convert"},
		{name: "ToUintptrArray - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintptrArray(tt.args)
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

func Test_ToFloat32Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToFloat32Array - []float32{float32(1)}", args: []float32{float32(1)}, want: [1]float32{float32(1)}},
		{name: "ToFloat32Array - []interface{}{float32(1)}", args: []interface{}{float32(1)}, want: [1]float32{float32(1)}},
		{name: "ToFloat32Array - [2]float32{float32(1), float32(1)}", args: [2]float32{float32(1), float32(1)}, want: [2]float32{float32(1), float32(1)}},
		{name: "ToFloat32Array - [2]interface{}{float32(1), float32(1)}", args: [2]interface{}{float32(1), float32(1)}, want: [2]float32{float32(1), float32(1)}},
		{name: "ToFloat32Array - str", args: "str", errContain: "can't convert"},
		{name: "ToFloat32Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32Array(tt.args)
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

func Test_ToFloat64Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToFloat64Array - []float64{float64(1)}", args: []float64{float64(1)}, want: [1]float64{float64(1)}},
		{name: "ToFloat64Array - []interface{}{float64(1)}", args: []interface{}{float64(1)}, want: [1]float64{float64(1)}},
		{name: "ToFloat64Array - [2]float64{float64(1), float64(1)}", args: [2]float64{float64(1), float64(1)}, want: [2]float64{float64(1), float64(1)}},
		{name: "ToFloat64Array - [2]interface{}{float64(1), float64(1)}", args: [2]interface{}{float64(1), float64(1)}, want: [2]float64{float64(1), float64(1)}},
		{name: "ToFloat64Array - str", args: "str", errContain: "can't convert"},
		{name: "ToFloat64Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64Array(tt.args)
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

func Test_ToBoolArray(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToBoolArray - []bool{bool(true)}", args: []bool{bool(true)}, want: [1]bool{bool(true)}},
		{name: "ToBoolArray - []interface{}{bool(true)}", args: []interface{}{bool(true)}, want: [1]bool{bool(true)}},
		{name: "ToBoolArray - [2]bool{bool(true), bool(true)}", args: [2]bool{bool(true), bool(true)}, want: [2]bool{bool(true), bool(true)}},
		{name: "ToBoolArray - [2]interface{}{bool(true), bool(true)}", args: [2]interface{}{bool(true), bool(true)}, want: [2]bool{bool(true), bool(true)}},
		{name: "ToBoolArray - str", args: "str", errContain: "can't convert"},
		{name: "ToBoolArray - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolArray(tt.args)
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

func Test_ToComplex64Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToComplex64Array - []complex64{complex64(1)}", args: []complex64{complex64(1)}, want: [1]complex64{complex64(1)}},
		{name: "ToComplex64Array - []interface{}{complex64(1)}", args: []interface{}{complex64(1)}, want: [1]complex64{complex64(1)}},
		{name: "ToComplex64Array - [2]complex64{complex64(1), complex64(1)}", args: [2]complex64{complex64(1), complex64(1)}, want: [2]complex64{complex64(1), complex64(1)}},
		{name: "ToComplex64Array - [2]interface{}{complex64(1), complex64(1)}", args: [2]interface{}{complex64(1), complex64(1)}, want: [2]complex64{complex64(1), complex64(1)}},
		{name: "ToComplex64Array - str", args: "str", errContain: "can't convert"},
		{name: "ToComplex64Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex64Array(tt.args)
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

func Test_ToComplex128Array(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToComplex128Array - []complex128{complex128(1)}", args: []complex128{complex128(1)}, want: [1]complex128{complex128(1)}},
		{name: "ToComplex128Array - []interface{}{complex128(1)}", args: []interface{}{complex128(1)}, want: [1]complex128{complex128(1)}},
		{name: "ToComplex128Array - [2]complex128{complex128(1), complex128(1)}", args: [2]complex128{complex128(1), complex128(1)}, want: [2]complex128{complex128(1), complex128(1)}},
		{name: "ToComplex128Array - [2]interface{}{complex128(1), complex128(1)}", args: [2]interface{}{complex128(1), complex128(1)}, want: [2]complex128{complex128(1), complex128(1)}},
		{name: "ToComplex128Array - str", args: "str", errContain: "can't convert"},
		{name: "ToComplex128Array - []string{-str-}", args: []string{"str"}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComplex128Array(tt.args)
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

func Test_ToStringArray(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       interface{}
		errContain string
	}{
		{name: "ToStringArray - []string{-1-}", args: []string{"1"}, want: [1]string{"1"}},
		{name: "ToStringArray - []interface{}{-1-}", args: []interface{}{"1"}, want: [1]string{"1"}},
		{name: "ToStringArray - [2]string{-1-, -2-}", args: [2]string{"1", "2"}, want: [2]string{"1", "2"}},
		{name: "ToStringArray - [2]interface{}{-1-, -2-}", args: [2]interface{}{"1", "2"}, want: [2]string{"1", "2"}},
		{name: "ToStringArray - str", args: "str", errContain: "can't convert"},
		{name: "ToStringArray - []int{1}", args: []int{1}, errContain: "can't convert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringArray(tt.args)
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
