package internal

import (
	"fmt"
)

func ToString(v interface{}) (string, error) {
	switch v := v.(type) {
	case []rune:
		return string(v), nil
	case []byte:
		return string(v), nil
	case string:
		return string(v), nil
	default:
		return "", fmt.Errorf("%v(%T) can't convert string", v, v)
	}
}
