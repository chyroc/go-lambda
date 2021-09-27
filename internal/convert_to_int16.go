package internal

import (
	"fmt"
)

func ToInt16(v interface{}) (int16, error) {
	switch v := v.(type) {
	case int:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to int16", v, v)
	}
}
