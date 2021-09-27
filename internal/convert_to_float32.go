package internal

import (
	"fmt"
)

func ToFloat32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case float32:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to float32", v, v)
	}
}
