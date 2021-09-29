package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

type itemArray struct{}

func Test_ToBasicArray(t *testing.T) {
	anyVal := itemArray{}
	as := assert.New(t)

	t.Run("ToIntArray", func(t *testing.T) {
		t.Run("success - []int{int(1)}", func(t *testing.T) {
			res, err := lambda.New([]int{int(1)}).ToIntArray()
			as.Nil(err)
			as.Equal([1]int{int(1)}, res)
		})

		t.Run("success - []interface{}{int(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int(1)}).ToIntArray()
			as.Nil(err)
			as.Equal([1]int{int(1)}, res)
		})

		t.Run("success - [2]int{int(1), int(2)}", func(t *testing.T) {
			res, err := lambda.New([2]int{int(1), int(2)}).ToIntArray()
			as.Nil(err)
			as.Equal([2]int{int(1), int(2)}, res)
		})

		t.Run("success - [2]interface{}{int(1), int(2)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int(1), int(2)}).ToIntArray()
			as.Nil(err)
			as.Equal([2]int{int(1), int(2)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToIntArray()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToIntArray()
			as.NotNil(err)
		})
	})
	t.Run("ToInt8Array", func(t *testing.T) {
		t.Run("success - []int8{int8(1)}", func(t *testing.T) {
			res, err := lambda.New([]int8{int8(1)}).ToInt8Array()
			as.Nil(err)
			as.Equal([1]int8{int8(1)}, res)
		})

		t.Run("success - []interface{}{int8(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int8(1)}).ToInt8Array()
			as.Nil(err)
			as.Equal([1]int8{int8(1)}, res)
		})

		t.Run("success - [2]int8{int8(1), int8(2)}", func(t *testing.T) {
			res, err := lambda.New([2]int8{int8(1), int8(2)}).ToInt8Array()
			as.Nil(err)
			as.Equal([2]int8{int8(1), int8(2)}, res)
		})

		t.Run("success - [2]interface{}{int8(1), int8(2)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int8(1), int8(2)}).ToInt8Array()
			as.Nil(err)
			as.Equal([2]int8{int8(1), int8(2)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt8Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt8Array()
			as.NotNil(err)
		})
	})
	t.Run("ToInt16Array", func(t *testing.T) {
		t.Run("success - []int16{int16(1)}", func(t *testing.T) {
			res, err := lambda.New([]int16{int16(1)}).ToInt16Array()
			as.Nil(err)
			as.Equal([1]int16{int16(1)}, res)
		})

		t.Run("success - []interface{}{int16(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int16(1)}).ToInt16Array()
			as.Nil(err)
			as.Equal([1]int16{int16(1)}, res)
		})

		t.Run("success - [2]int16{int16(1), int16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int16{int16(1), int16(1)}).ToInt16Array()
			as.Nil(err)
			as.Equal([2]int16{int16(1), int16(1)}, res)
		})

		t.Run("success - [2]interface{}{int16(1), int16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int16(1), int16(1)}).ToInt16Array()
			as.Nil(err)
			as.Equal([2]int16{int16(1), int16(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt16Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt16Array()
			as.NotNil(err)
		})
	})
	t.Run("ToInt32Array", func(t *testing.T) {
		t.Run("success - []int32{int32(1)}", func(t *testing.T) {
			res, err := lambda.New([]int32{int32(1)}).ToInt32Array()
			as.Nil(err)
			as.Equal([1]int32{int32(1)}, res)
		})

		t.Run("success - []interface{}{int32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int32(1)}).ToInt32Array()
			as.Nil(err)
			as.Equal([1]int32{int32(1)}, res)
		})

		t.Run("success - [2]int32{int32(1), int32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int32{int32(1), int32(1)}).ToInt32Array()
			as.Nil(err)
			as.Equal([2]int32{int32(1), int32(1)}, res)
		})

		t.Run("success - [2]interface{}{int32(1), int32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int32(1), int32(1)}).ToInt32Array()
			as.Nil(err)
			as.Equal([2]int32{int32(1), int32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt32Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt32Array()
			as.NotNil(err)
		})
	})
	t.Run("ToInt64Array", func(t *testing.T) {
		t.Run("success - []int64{int64(1)}", func(t *testing.T) {
			res, err := lambda.New([]int64{int64(1)}).ToInt64Array()
			as.Nil(err)
			as.Equal([1]int64{int64(1)}, res)
		})

		t.Run("success - []interface{}{int64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{int64(1)}).ToInt64Array()
			as.Nil(err)
			as.Equal([1]int64{int64(1)}, res)
		})

		t.Run("success - [2]int64{int64(1), int64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]int64{int64(1), int64(1)}).ToInt64Array()
			as.Nil(err)
			as.Equal([2]int64{int64(1), int64(1)}, res)
		})

		t.Run("success - [2]interface{}{int64(1), int64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{int64(1), int64(1)}).ToInt64Array()
			as.Nil(err)
			as.Equal([2]int64{int64(1), int64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt64Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt64Array()
			as.NotNil(err)
		})
	})
	t.Run("ToUintArray", func(t *testing.T) {
		t.Run("success - []uint{uint(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint{uint(1)}).ToUintArray()
			as.Nil(err)
			as.Equal([1]uint{uint(1)}, res)
		})

		t.Run("success - []interface{}{uint(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint(1)}).ToUintArray()
			as.Nil(err)
			as.Equal([1]uint{uint(1)}, res)
		})

		t.Run("success - [2]uint{uint(1), uint(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint{uint(1), uint(1)}).ToUintArray()
			as.Nil(err)
			as.Equal([2]uint{uint(1), uint(1)}, res)
		})

		t.Run("success - [2]interface{}{uint(1), uint(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint(1), uint(1)}).ToUintArray()
			as.Nil(err)
			as.Equal([2]uint{uint(1), uint(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUintArray()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUintArray()
			as.NotNil(err)
		})
	})
	t.Run("ToUint8Array", func(t *testing.T) {
		t.Run("success - []uint8{uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint8{uint8(1)}).ToUint8Array()
			as.Nil(err)
			as.Equal([1]uint8{uint8(1)}, res)
		})

		t.Run("success - []interface{}{uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint8(1)}).ToUint8Array()
			as.Nil(err)
			as.Equal([1]uint8{uint8(1)}, res)
		})

		t.Run("success - [2]uint8{uint8(1), uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint8{uint8(1), uint8(1)}).ToUint8Array()
			as.Nil(err)
			as.Equal([2]uint8{uint8(1), uint8(1)}, res)
		})

		t.Run("success - [2]interface{}{uint8(1), uint8(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint8(1), uint8(1)}).ToUint8Array()
			as.Nil(err)
			as.Equal([2]uint8{uint8(1), uint8(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint8Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint8Array()
			as.NotNil(err)
		})
	})
	t.Run("ToUint16Array", func(t *testing.T) {
		t.Run("success - []uint16{uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint16{uint16(1)}).ToUint16Array()
			as.Nil(err)
			as.Equal([1]uint16{uint16(1)}, res)
		})

		t.Run("success - []interface{}{uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint16(1)}).ToUint16Array()
			as.Nil(err)
			as.Equal([1]uint16{uint16(1)}, res)
		})

		t.Run("success - [2]uint16{uint16(1), uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint16{uint16(1), uint16(1)}).ToUint16Array()
			as.Nil(err)
			as.Equal([2]uint16{uint16(1), uint16(1)}, res)
		})

		t.Run("success - [2]interface{}{uint16(1), uint16(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint16(1), uint16(1)}).ToUint16Array()
			as.Nil(err)
			as.Equal([2]uint16{uint16(1), uint16(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint16Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint16Array()
			as.NotNil(err)
		})
	})
	t.Run("ToUint32Array", func(t *testing.T) {
		t.Run("success - []uint32{uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint32{uint32(1)}).ToUint32Array()
			as.Nil(err)
			as.Equal([1]uint32{uint32(1)}, res)
		})

		t.Run("success - []interface{}{uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint32(1)}).ToUint32Array()
			as.Nil(err)
			as.Equal([1]uint32{uint32(1)}, res)
		})

		t.Run("success - [2]uint32{uint32(1), uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint32{uint32(1), uint32(1)}).ToUint32Array()
			as.Nil(err)
			as.Equal([2]uint32{uint32(1), uint32(1)}, res)
		})

		t.Run("success - [2]interface{}{uint32(1), uint32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint32(1), uint32(1)}).ToUint32Array()
			as.Nil(err)
			as.Equal([2]uint32{uint32(1), uint32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint32Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint32Array()
			as.NotNil(err)
		})
	})
	t.Run("ToUint64Array", func(t *testing.T) {
		t.Run("success - []uint64{uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([]uint64{uint64(1)}).ToUint64Array()
			as.Nil(err)
			as.Equal([1]uint64{uint64(1)}, res)
		})

		t.Run("success - []interface{}{uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uint64(1)}).ToUint64Array()
			as.Nil(err)
			as.Equal([1]uint64{uint64(1)}, res)
		})

		t.Run("success - [2]uint64{uint64(1), uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uint64{uint64(1), uint64(1)}).ToUint64Array()
			as.Nil(err)
			as.Equal([2]uint64{uint64(1), uint64(1)}, res)
		})

		t.Run("success - [2]interface{}{uint64(1), uint64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uint64(1), uint64(1)}).ToUint64Array()
			as.Nil(err)
			as.Equal([2]uint64{uint64(1), uint64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint64Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint64Array()
			as.NotNil(err)
		})
	})
	t.Run("ToUintptrArray", func(t *testing.T) {
		t.Run("success - []uintptr{uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([]uintptr{uintptr(1)}).ToUintptrArray()
			as.Nil(err)
			as.Equal([1]uintptr{uintptr(1)}, res)
		})

		t.Run("success - []interface{}{uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{uintptr(1)}).ToUintptrArray()
			as.Nil(err)
			as.Equal([1]uintptr{uintptr(1)}, res)
		})

		t.Run("success - [2]uintptr{uintptr(1), uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([2]uintptr{uintptr(1), uintptr(1)}).ToUintptrArray()
			as.Nil(err)
			as.Equal([2]uintptr{uintptr(1), uintptr(1)}, res)
		})

		t.Run("success - [2]interface{}{uintptr(1), uintptr(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{uintptr(1), uintptr(1)}).ToUintptrArray()
			as.Nil(err)
			as.Equal([2]uintptr{uintptr(1), uintptr(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUintptrArray()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUintptrArray()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat32Array", func(t *testing.T) {
		t.Run("success - []float32{float32(1)}", func(t *testing.T) {
			res, err := lambda.New([]float32{float32(1)}).ToFloat32Array()
			as.Nil(err)
			as.Equal([1]float32{float32(1)}, res)
		})

		t.Run("success - []interface{}{float32(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{float32(1)}).ToFloat32Array()
			as.Nil(err)
			as.Equal([1]float32{float32(1)}, res)
		})

		t.Run("success - [2]float32{float32(1), float32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]float32{float32(1), float32(1)}).ToFloat32Array()
			as.Nil(err)
			as.Equal([2]float32{float32(1), float32(1)}, res)
		})

		t.Run("success - [2]interface{}{float32(1), float32(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{float32(1), float32(1)}).ToFloat32Array()
			as.Nil(err)
			as.Equal([2]float32{float32(1), float32(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat32Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat32Array()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat64Array", func(t *testing.T) {
		t.Run("success - []float64{float64(1)}", func(t *testing.T) {
			res, err := lambda.New([]float64{float64(1)}).ToFloat64Array()
			as.Nil(err)
			as.Equal([1]float64{float64(1)}, res)
		})

		t.Run("success - []interface{}{float64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{float64(1)}).ToFloat64Array()
			as.Nil(err)
			as.Equal([1]float64{float64(1)}, res)
		})

		t.Run("success - [2]float64{float64(1), float64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]float64{float64(1), float64(1)}).ToFloat64Array()
			as.Nil(err)
			as.Equal([2]float64{float64(1), float64(1)}, res)
		})

		t.Run("success - [2]interface{}{float64(1), float64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{float64(1), float64(1)}).ToFloat64Array()
			as.Nil(err)
			as.Equal([2]float64{float64(1), float64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat64Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat64Array()
			as.NotNil(err)
		})
	})
	t.Run("ToBoolArray", func(t *testing.T) {
		t.Run("success - []bool{bool(true)}", func(t *testing.T) {
			res, err := lambda.New([]bool{bool(true)}).ToBoolArray()
			as.Nil(err)
			as.Equal([1]bool{bool(true)}, res)
		})

		t.Run("success - []interface{}{bool(true)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{bool(true)}).ToBoolArray()
			as.Nil(err)
			as.Equal([1]bool{bool(true)}, res)
		})

		t.Run("success - [2]bool{bool(true), bool(true)}", func(t *testing.T) {
			res, err := lambda.New([2]bool{bool(true), bool(true)}).ToBoolArray()
			as.Nil(err)
			as.Equal([2]bool{bool(true), bool(true)}, res)
		})

		t.Run("success - [2]interface{}{bool(true), bool(true)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{bool(true), bool(true)}).ToBoolArray()
			as.Nil(err)
			as.Equal([2]bool{bool(true), bool(true)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToBoolArray()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToBoolArray()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex64Array", func(t *testing.T) {
		t.Run("success - []complex64{complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([]complex64{complex64(1)}).ToComplex64Array()
			as.Nil(err)
			as.Equal([1]complex64{complex64(1)}, res)
		})

		t.Run("success - []interface{}{complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{complex64(1)}).ToComplex64Array()
			as.Nil(err)
			as.Equal([1]complex64{complex64(1)}, res)
		})

		t.Run("success - [2]complex64{complex64(1), complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]complex64{complex64(1), complex64(1)}).ToComplex64Array()
			as.Nil(err)
			as.Equal([2]complex64{complex64(1), complex64(1)}, res)
		})

		t.Run("success - [2]interface{}{complex64(1), complex64(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{complex64(1), complex64(1)}).ToComplex64Array()
			as.Nil(err)
			as.Equal([2]complex64{complex64(1), complex64(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex64Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex64Array()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex128Array", func(t *testing.T) {
		t.Run("success - []complex128{complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([]complex128{complex128(1)}).ToComplex128Array()
			as.Nil(err)
			as.Equal([1]complex128{complex128(1)}, res)
		})

		t.Run("success - []interface{}{complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{complex128(1)}).ToComplex128Array()
			as.Nil(err)
			as.Equal([1]complex128{complex128(1)}, res)
		})

		t.Run("success - [2]complex128{complex128(1), complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([2]complex128{complex128(1), complex128(1)}).ToComplex128Array()
			as.Nil(err)
			as.Equal([2]complex128{complex128(1), complex128(1)}, res)
		})

		t.Run("success - [2]interface{}{complex128(1), complex128(1)}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{complex128(1), complex128(1)}).ToComplex128Array()
			as.Nil(err)
			as.Equal([2]complex128{complex128(1), complex128(1)}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex128Array()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex128Array()
			as.NotNil(err)
		})
	})
	t.Run("ToStringArray", func(t *testing.T) {
		t.Run("success - []string{-1-}", func(t *testing.T) {
			res, err := lambda.New([]string{"1"}).ToStringArray()
			as.Nil(err)
			as.Equal([1]string{"1"}, res)
		})

		t.Run("success - []interface{}{-1-}", func(t *testing.T) {
			res, err := lambda.New([]interface{}{"1"}).ToStringArray()
			as.Nil(err)
			as.Equal([1]string{"1"}, res)
		})

		t.Run("success - [2]string{-1-, -2-}", func(t *testing.T) {
			res, err := lambda.New([2]string{"1", "2"}).ToStringArray()
			as.Nil(err)
			as.Equal([2]string{"1", "2"}, res)
		})

		t.Run("success - [2]interface{}{-1-, -2-}", func(t *testing.T) {
			res, err := lambda.New([2]interface{}{"1", "2"}).ToStringArray()
			as.Nil(err)
			as.Equal([2]string{"1", "2"}, res)
		})

		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToStringArray()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToStringArray()
			as.NotNil(err)
		})
	})
}
