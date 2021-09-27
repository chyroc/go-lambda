package internal

import (
	"fmt"
)

func ToUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case uint:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) can't convert to uint", v, v)
	}
}
