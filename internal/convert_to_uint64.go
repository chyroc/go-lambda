package internal

import (
	"fmt"
)

func ToUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to uint64", v, v)
	}
}
