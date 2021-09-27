package internal

import (
	"fmt"
)

func ToInt8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case int:
		return int8(v), nil
	case int8:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to int8", v, v)
	}
}
