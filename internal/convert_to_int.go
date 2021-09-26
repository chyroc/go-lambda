package internal

import (
	"fmt"
)

func ToInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case int:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to int", v, v)
	}
}
