package internal

import (
	"fmt"
)

func ToInt32(v interface{}) (int32, error) {
	switch v := v.(type) {
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to int32", v, v)
	}
}
