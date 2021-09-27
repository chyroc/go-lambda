package internal

import (
	"reflect"
)

func IsZero(v interface{}) bool {
	// fmt.Printf("=%v, %T=", v, v)
	if v == nil {
		return true
	}
	switch v := v.(type) {
	case error:
		return v == nil
	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32: // rune
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case uint8: // byte
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	case uintptr:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return !v
	case complex64:
		return v == 0
	case complex128:
		return v == 0
	case interface{}:
		vv := reflect.ValueOf(v)
		if ReflectCanInterface(vv) && vv.IsNil() {
			return true
		}
		return v == nil
	}

	return v == nil
}
