package internal

import (
	"fmt"
)

func ToUint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case uint:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to uint32", v, v)
	}
}
