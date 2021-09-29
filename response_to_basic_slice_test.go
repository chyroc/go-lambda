package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

type itemSlice struct{}

func Test_ToBasicSlice(t *testing.T) {
	anyVal := itemSlice{}
	as := assert.New(t)

	t.Run("ToIntSlice", func(t *testing.T) {
		t.Run("success - []int{int(1)}", func(t *testing.T) {
			res, err := lambda.New([]int{int(1)}).ToIntSlice()
			as.Nil(err)
			as.Equal([]int{int(1)}, res)
		})

		t.Run("success - []interface{}{int(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int(1)}).ToIntSlice()
			as.Nil(err)
			as.Equal([]int{int(1)}, res)
		})

		t.Run("success - [2]int{int(1), int(2)}", func(t *testing.T) {
			res, err := lambda.New([2]int{int(1), int(2)}).ToIntSlice()
			as.Nil(err)
			as.Equal([]int{int(1), int(2)}, res)
		})

		t.Run("success - [2]interface{}{int(1), int(2)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int(1), int(2)}).ToIntSlice()
			as.Nil(err)
			as.Equal([]int{int(1), int(2)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToIntSlice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToIntSlice()
			as.NotNil(err)
		})
	})
	t.Run("ToInt8Slice", func(t *testing.T) {
		t.Run("success - []int8{int8(1)}", func(t *testing.T) {
			res, err := lambda.New([]int8{int8(1)}).ToInt8Slice()
			as.Nil(err)
			as.Equal([]int8{int8(1)}, res)
		})

		t.Run("success - []interface{}{int8(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int8(1)}).ToInt8Slice()
			as.Nil(err)
			as.Equal([]int8{int8(1)}, res)
		})

		t.Run("success - [2]int8{int8(1), int8(2)}", func(t *testing.T) {
			res, err := lambda.New([2]int8{int8(1), int8(2)}).ToInt8Slice()
			as.Nil(err)
			as.Equal([]int8{int8(1), int8(2)}, res)
		})

		t.Run("success - [2]interface{}{int8(1), int8(2)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int8(1), int8(2)}).ToInt8Slice()
			as.Nil(err)
			as.Equal([]int8{int8(1), int8(2)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt8Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt8Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToInt16Slice", func(t *testing.T) {
		t.Run("success - []int16{int16(1)}", func(t *testing.T) {
			res, err := lambda.New([]int16{int16(1)}).ToInt16Slice()
			as.Nil(err)
			as.Equal([]int16{int16(1)}, res)
		})

		t.Run("success - []interface{}{int16(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int16(1)}).ToInt16Slice()
			as.Nil(err)
			as.Equal([]int16{int16(1)}, res)
		})

		t.Run("success - [2]int16{int16(1), int16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int16{int16(1), int16(1)}).ToInt16Slice()
			as.Nil(err)
			as.Equal([]int16{int16(1), int16(1)}, res)
		})

		t.Run("success - [2]interface{}{int16(1), int16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int16(1), int16(1)}).ToInt16Slice()
			as.Nil(err)
			as.Equal([]int16{int16(1), int16(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt16Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt16Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToInt32Slice", func(t *testing.T) {
		t.Run("success - []int32{int32(1)}", func(t *testing.T) {
			res, err := lambda.New([]int32{int32(1)}).ToInt32Slice()
			as.Nil(err)
			as.Equal([]int32{int32(1)}, res)
		})

		t.Run("success - []interface{}{int32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int32(1)}).ToInt32Slice()
			as.Nil(err)
			as.Equal([]int32{int32(1)}, res)
		})

		t.Run("success - [2]int32{int32(1), int32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int32{int32(1), int32(1)}).ToInt32Slice()
			as.Nil(err)
			as.Equal([]int32{int32(1), int32(1)}, res)
		})

		t.Run("success - [2]interface{}{int32(1), int32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int32(1), int32(1)}).ToInt32Slice()
			as.Nil(err)
			as.Equal([]int32{int32(1), int32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt32Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt32Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToInt64Slice", func(t *testing.T) {
		t.Run("success - []int64{int64(1)}", func(t *testing.T) {
			res, err := lambda.New([]int64{int64(1)}).ToInt64Slice()
			as.Nil(err)
			as.Equal([]int64{int64(1)}, res)
		})

		t.Run("success - []interface{}{int64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int64(1)}).ToInt64Slice()
			as.Nil(err)
			as.Equal([]int64{int64(1)}, res)
		})

		t.Run("success - [2]int64{int64(1), int64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int64{int64(1), int64(1)}).ToInt64Slice()
			as.Nil(err)
			as.Equal([]int64{int64(1), int64(1)}, res)
		})

		t.Run("success - [2]interface{}{int64(1), int64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int64(1), int64(1)}).ToInt64Slice()
			as.Nil(err)
			as.Equal([]int64{int64(1), int64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt64Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt64Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToUintSlice", func(t *testing.T) {
		t.Run("success - []uint{uint(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint{uint(1)}).ToUintSlice()
			as.Nil(err)
			as.Equal([]uint{uint(1)}, res)
		})

		t.Run("success - []interface{}{uint(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint(1)}).ToUintSlice()
			as.Nil(err)
			as.Equal([]uint{uint(1)}, res)
		})

		t.Run("success - [2]uint{uint(1), uint(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint{uint(1), uint(1)}).ToUintSlice()
			as.Nil(err)
			as.Equal([]uint{uint(1), uint(1)}, res)
		})

		t.Run("success - [2]interface{}{uint(1), uint(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint(1), uint(1)}).ToUintSlice()
			as.Nil(err)
			as.Equal([]uint{uint(1), uint(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUintSlice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUintSlice()
			as.NotNil(err)
		})
	})
	t.Run("ToUint8Slice", func(t *testing.T) {
		t.Run("success - []uint8{uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint8{uint8(1)}).ToUint8Slice()
			as.Nil(err)
			as.Equal([]uint8{uint8(1)}, res)
		})

		t.Run("success - []interface{}{uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint8(1)}).ToUint8Slice()
			as.Nil(err)
			as.Equal([]uint8{uint8(1)}, res)
		})

		t.Run("success - [2]uint8{uint8(1), uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint8{uint8(1), uint8(1)}).ToUint8Slice()
			as.Nil(err)
			as.Equal([]uint8{uint8(1), uint8(1)}, res)
		})

		t.Run("success - [2]interface{}{uint8(1), uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint8(1), uint8(1)}).ToUint8Slice()
			as.Nil(err)
			as.Equal([]uint8{uint8(1), uint8(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint8Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint8Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToUint16Slice", func(t *testing.T) {
		t.Run("success - []uint16{uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint16{uint16(1)}).ToUint16Slice()
			as.Nil(err)
			as.Equal([]uint16{uint16(1)}, res)
		})

		t.Run("success - []interface{}{uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint16(1)}).ToUint16Slice()
			as.Nil(err)
			as.Equal([]uint16{uint16(1)}, res)
		})

		t.Run("success - [2]uint16{uint16(1), uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint16{uint16(1), uint16(1)}).ToUint16Slice()
			as.Nil(err)
			as.Equal([]uint16{uint16(1), uint16(1)}, res)
		})

		t.Run("success - [2]interface{}{uint16(1), uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint16(1), uint16(1)}).ToUint16Slice()
			as.Nil(err)
			as.Equal([]uint16{uint16(1), uint16(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint16Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint16Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToUint32Slice", func(t *testing.T) {
		t.Run("success - []uint32{uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint32{uint32(1)}).ToUint32Slice()
			as.Nil(err)
			as.Equal([]uint32{uint32(1)}, res)
		})

		t.Run("success - []interface{}{uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint32(1)}).ToUint32Slice()
			as.Nil(err)
			as.Equal([]uint32{uint32(1)}, res)
		})

		t.Run("success - [2]uint32{uint32(1), uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint32{uint32(1), uint32(1)}).ToUint32Slice()
			as.Nil(err)
			as.Equal([]uint32{uint32(1), uint32(1)}, res)
		})

		t.Run("success - [2]interface{}{uint32(1), uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint32(1), uint32(1)}).ToUint32Slice()
			as.Nil(err)
			as.Equal([]uint32{uint32(1), uint32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint32Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint32Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToUint64Slice", func(t *testing.T) {
		t.Run("success - []uint64{uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint64{uint64(1)}).ToUint64Slice()
			as.Nil(err)
			as.Equal([]uint64{uint64(1)}, res)
		})

		t.Run("success - []interface{}{uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint64(1)}).ToUint64Slice()
			as.Nil(err)
			as.Equal([]uint64{uint64(1)}, res)
		})

		t.Run("success - [2]uint64{uint64(1), uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint64{uint64(1), uint64(1)}).ToUint64Slice()
			as.Nil(err)
			as.Equal([]uint64{uint64(1), uint64(1)}, res)
		})

		t.Run("success - [2]interface{}{uint64(1), uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint64(1), uint64(1)}).ToUint64Slice()
			as.Nil(err)
			as.Equal([]uint64{uint64(1), uint64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint64Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint64Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToUintptrSlice", func(t *testing.T) {
		t.Run("success - []uintptr{uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([]uintptr{uintptr(1)}).ToUintptrSlice()
			as.Nil(err)
			as.Equal([]uintptr{uintptr(1)}, res)
		})

		t.Run("success - []interface{}{uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uintptr(1)}).ToUintptrSlice()
			as.Nil(err)
			as.Equal([]uintptr{uintptr(1)}, res)
		})

		t.Run("success - [2]uintptr{uintptr(1), uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uintptr{uintptr(1), uintptr(1)}).ToUintptrSlice()
			as.Nil(err)
			as.Equal([]uintptr{uintptr(1), uintptr(1)}, res)
		})

		t.Run("success - [2]interface{}{uintptr(1), uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uintptr(1), uintptr(1)}).ToUintptrSlice()
			as.Nil(err)
			as.Equal([]uintptr{uintptr(1), uintptr(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUintptrSlice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUintptrSlice()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat32Slice", func(t *testing.T) {
		t.Run("success - []float32{float32(1)}", func(t *testing.T) {
			res, err := lambda.New([]float32{float32(1)}).ToFloat32Slice()
			as.Nil(err)
			as.Equal([]float32{float32(1)}, res)
		})

		t.Run("success - []interface{}{float32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{float32(1)}).ToFloat32Slice()
			as.Nil(err)
			as.Equal([]float32{float32(1)}, res)
		})

		t.Run("success - [2]float32{float32(1), float32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]float32{float32(1), float32(1)}).ToFloat32Slice()
			as.Nil(err)
			as.Equal([]float32{float32(1), float32(1)}, res)
		})

		t.Run("success - [2]interface{}{float32(1), float32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{float32(1), float32(1)}).ToFloat32Slice()
			as.Nil(err)
			as.Equal([]float32{float32(1), float32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat32Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat32Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat64Slice", func(t *testing.T) {
		t.Run("success - []float64{float64(1)}", func(t *testing.T) {
			res, err := lambda.New([]float64{float64(1)}).ToFloat64Slice()
			as.Nil(err)
			as.Equal([]float64{float64(1)}, res)
		})

		t.Run("success - []interface{}{float64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{float64(1)}).ToFloat64Slice()
			as.Nil(err)
			as.Equal([]float64{float64(1)}, res)
		})

		t.Run("success - [2]float64{float64(1), float64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]float64{float64(1), float64(1)}).ToFloat64Slice()
			as.Nil(err)
			as.Equal([]float64{float64(1), float64(1)}, res)
		})

		t.Run("success - [2]interface{}{float64(1), float64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{float64(1), float64(1)}).ToFloat64Slice()
			as.Nil(err)
			as.Equal([]float64{float64(1), float64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat64Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat64Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToBoolSlice", func(t *testing.T) {
		t.Run("success - []bool{bool(true)}", func(t *testing.T) {
			res, err := lambda.New([]bool{bool(true)}).ToBoolSlice()
			as.Nil(err)
			as.Equal([]bool{bool(true)}, res)
		})

		t.Run("success - []interface{}{bool(true)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{bool(true)}).ToBoolSlice()
			as.Nil(err)
			as.Equal([]bool{bool(true)}, res)
		})

		t.Run("success - [2]bool{bool(true), bool(true)}", func(t *testing.T) {
			res, err := lambda.New([2]bool{bool(true), bool(true)}).ToBoolSlice()
			as.Nil(err)
			as.Equal([]bool{bool(true), bool(true)}, res)
		})

		t.Run("success - [2]interface{}{bool(true), bool(true)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{bool(true), bool(true)}).ToBoolSlice()
			as.Nil(err)
			as.Equal([]bool{bool(true), bool(true)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToBoolSlice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToBoolSlice()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex64Slice", func(t *testing.T) {
		t.Run("success - []complex64{complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([]complex64{complex64(1)}).ToComplex64Slice()
			as.Nil(err)
			as.Equal([]complex64{complex64(1)}, res)
		})

		t.Run("success - []interface{}{complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{complex64(1)}).ToComplex64Slice()
			as.Nil(err)
			as.Equal([]complex64{complex64(1)}, res)
		})

		t.Run("success - [2]complex64{complex64(1), complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]complex64{complex64(1), complex64(1)}).ToComplex64Slice()
			as.Nil(err)
			as.Equal([]complex64{complex64(1), complex64(1)}, res)
		})

		t.Run("success - [2]interface{}{complex64(1), complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{complex64(1), complex64(1)}).ToComplex64Slice()
			as.Nil(err)
			as.Equal([]complex64{complex64(1), complex64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex64Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex64Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex128Slice", func(t *testing.T) {
		t.Run("success - []complex128{complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([]complex128{complex128(1)}).ToComplex128Slice()
			as.Nil(err)
			as.Equal([]complex128{complex128(1)}, res)
		})

		t.Run("success - []interface{}{complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{complex128(1)}).ToComplex128Slice()
			as.Nil(err)
			as.Equal([]complex128{complex128(1)}, res)
		})

		t.Run("success - [2]complex128{complex128(1), complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([2]complex128{complex128(1), complex128(1)}).ToComplex128Slice()
			as.Nil(err)
			as.Equal([]complex128{complex128(1), complex128(1)}, res)
		})

		t.Run("success - [2]interface{}{complex128(1), complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{complex128(1), complex128(1)}).ToComplex128Slice()
			as.Nil(err)
			as.Equal([]complex128{complex128(1), complex128(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex128Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex128Slice()
			as.NotNil(err)
		})
	})
	t.Run("ToStringSlice", func(t *testing.T) {
		t.Run("success - []string{-1-}", func(t *testing.T) {
			res, err := lambda.New([]string{"1"}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1"}, res)
		})

		t.Run("success - []interface{}{-1-}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{"1"}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1"}, res)
		})

		t.Run("success - [2]string{-1-, -2-}", func(t *testing.T) {
			res, err := lambda.New([2]string{"1", "2"}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1", "2"}, res)
		})

		t.Run("success - [2]interface{}{-1-, -2-}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{"1", "2"}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1", "2"}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToStringSlice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToStringSlice()
			as.NotNil(err)
		})
	})
}
