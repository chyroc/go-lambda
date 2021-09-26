package internal

import (
	"fmt"
)

func ToFloat(v interface{}) (float64, error) {
	switch v := v.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to float64", v, v)
	}
}
