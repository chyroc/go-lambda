package internal

import (
	"fmt"
)

func ToUint16(v interface{}) (uint16, error) {
	switch v := v.(type) {
	case uint:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to uint16", v, v)
	}
}
