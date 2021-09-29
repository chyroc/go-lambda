package internal

import (
	"fmt"
	"reflect"
)

func ToIntSlice(v interface{}) (resp []int, err error) {
	switch v := v.(type) {
	case []int:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				ii, err := ToInt(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				resp = append(resp, ii)
			}
			if !catchErr {
				return resp, nil
			}
		}
		return nil, fmt.Errorf("%v(%T) can't convert to []int", v, v)
	}
}

func ToInt8Slice(v interface{}) (resp []int8, err error) {
	switch v := v.(type) {
	case []int8:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt8(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []int8", v, v)
	}
}

func ToInt16Slice(v interface{}) (resp []int16, err error) {
	switch v := v.(type) {
	case []int16:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt16(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []int16", v, v)
	}
}

func ToInt32Slice(v interface{}) (resp []int32, err error) {
	switch v := v.(type) {
	case []int32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt32(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []int32", v, v)
	}
}

func ToInt64Slice(v interface{}) (resp []int64, err error) {
	switch v := v.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToInt64(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []int64", v, v)
	}
}

func ToUintSlice(v interface{}) (resp []uint, err error) {
	switch v := v.(type) {
	case []uint:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []uint", v, v)
	}
}

func ToUint8Slice(v interface{}) (resp []uint8, err error) {
	switch v := v.(type) {
	case []uint8:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint8(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []uint8", v, v)
	}
}

func ToUint16Slice(v interface{}) (resp []uint16, err error) {
	switch v := v.(type) {
	case []uint16:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint16(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []uint16", v, v)
	}
}

func ToUint32Slice(v interface{}) (resp []uint32, err error) {
	switch v := v.(type) {
	case []uint32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint32(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []uint32", v, v)
	}
}

func ToUint64Slice(v interface{}) (resp []uint64, err error) {
	switch v := v.(type) {
	case []uint64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToUint64(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []uint64", v, v)
	}
}

func ToFloat32Slice(v interface{}) (resp []float32, err error) {
	switch v := v.(type) {
	case []float32:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToFloat32(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []float32", v, v)
	}
}

func ToFloat64Slice(v interface{}) (resp []float64, err error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToFloat64(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []float64", v, v)
	}
}

func ToBoolSlice(v interface{}) (resp []bool, err error) {
	switch v := v.(type) {
	case []bool:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToBool(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []bool", v, v)
	}
}

func ToComplex64Slice(v interface{}) (resp []complex64, err error) {
	switch v := v.(type) {
	case []complex64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToComplex64(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []complex64", v, v)
	}
}

func ToComplex128Slice(v interface{}) (resp []complex128, err error) {
	switch v := v.(type) {
	case []complex128:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToComplex128(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []complex128", v, v)
	}
}

func ToStringSlice(v interface{}) (resp []string, err error) {
	switch v := v.(type) {
	case []string:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToString(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []string", v, v)
	}
}
