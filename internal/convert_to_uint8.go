package internal

import (
	"fmt"
)

func ToUint8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case uint:
		return uint8(v), nil
	case uint8:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to uint8", v, v)
	}
}
