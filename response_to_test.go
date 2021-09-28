package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_To(t *testing.T) {
	anyVal := item{}
	as := assert.New(t)

	t.Run("ToInt", func(t *testing.T) {
		t.Run("ToInt - success", func(t *testing.T) {
			res, err := lambda.New(int(1)).ToInt()
			as.Nil(err)
			as.Equal(int(int(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt()
			as.NotNil(err)
		})
	})
	t.Run("ToInt8", func(t *testing.T) {
		t.Run("ToInt8 - success", func(t *testing.T) {
			res, err := lambda.New(int8(1)).ToInt8()
			as.Nil(err)
			as.Equal(int8(int8(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt8()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt8()
			as.NotNil(err)
		})
	})
	t.Run("ToInt16", func(t *testing.T) {
		t.Run("ToInt16 - success", func(t *testing.T) {
			res, err := lambda.New(int16(1)).ToInt16()
			as.Nil(err)
			as.Equal(int16(int16(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt16()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt16()
			as.NotNil(err)
		})
	})
	t.Run("ToInt32", func(t *testing.T) {
		t.Run("ToInt32 - success", func(t *testing.T) {
			res, err := lambda.New(int32(1)).ToInt32()
			as.Nil(err)
			as.Equal(int32(int32(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt32()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt32()
			as.NotNil(err)
		})
	})
	t.Run("ToInt64", func(t *testing.T) {
		t.Run("ToInt64 - success", func(t *testing.T) {
			res, err := lambda.New(int64(1)).ToInt64()
			as.Nil(err)
			as.Equal(int64(int64(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToInt64()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToInt64()
			as.NotNil(err)
		})
	})
	t.Run("ToUint", func(t *testing.T) {
		t.Run("ToUint - success", func(t *testing.T) {
			res, err := lambda.New(uint(1)).ToUint()
			as.Nil(err)
			as.Equal(uint(uint(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint()
			as.NotNil(err)
		})
	})
	t.Run("ToUint8", func(t *testing.T) {
		t.Run("ToUint8 - success", func(t *testing.T) {
			res, err := lambda.New(uint8(1)).ToUint8()
			as.Nil(err)
			as.Equal(uint8(uint8(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint8()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint8()
			as.NotNil(err)
		})
	})
	t.Run("ToUint16", func(t *testing.T) {
		t.Run("ToUint16 - success", func(t *testing.T) {
			res, err := lambda.New(uint16(1)).ToUint16()
			as.Nil(err)
			as.Equal(uint16(uint16(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint16()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint16()
			as.NotNil(err)
		})
	})
	t.Run("ToUint32", func(t *testing.T) {
		t.Run("ToUint32 - success", func(t *testing.T) {
			res, err := lambda.New(uint32(1)).ToUint32()
			as.Nil(err)
			as.Equal(uint32(uint32(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint32()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint32()
			as.NotNil(err)
		})
	})
	t.Run("ToUint64", func(t *testing.T) {
		t.Run("ToUint64 - success", func(t *testing.T) {
			res, err := lambda.New(uint64(1)).ToUint64()
			as.Nil(err)
			as.Equal(uint64(uint64(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToUint64()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToUint64()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat32", func(t *testing.T) {
		t.Run("ToFloat32 - success", func(t *testing.T) {
			res, err := lambda.New(float32(1)).ToFloat32()
			as.Nil(err)
			as.Equal(float32(float32(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat32()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat32()
			as.NotNil(err)
		})
	})
	t.Run("ToFloat64", func(t *testing.T) {
		t.Run("ToFloat64 - success", func(t *testing.T) {
			res, err := lambda.New(float64(1)).ToFloat64()
			as.Nil(err)
			as.Equal(float64(float64(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToFloat64()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToFloat64()
			as.NotNil(err)
		})
	})
	t.Run("ToBool", func(t *testing.T) {
		t.Run("ToBool - success", func(t *testing.T) {
			res, err := lambda.New(bool(1)).ToBool()
			as.Nil(err)
			as.Equal(bool(true), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToBool()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToBool()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex64", func(t *testing.T) {
		t.Run("ToComplex64 - success", func(t *testing.T) {
			res, err := lambda.New(complex64(1)).ToComplex64()
			as.Nil(err)
			as.Equal(complex64(complex64(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex64()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex64()
			as.NotNil(err)
		})
	})
	t.Run("ToComplex128", func(t *testing.T) {
		t.Run("ToComplex128 - success", func(t *testing.T) {
			res, err := lambda.New(complex128(1)).ToComplex128()
			as.Nil(err)
			as.Equal(complex128(complex128(1)), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).ToComplex128()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).ToComplex128()
			as.NotNil(err)
		})
	})
}
