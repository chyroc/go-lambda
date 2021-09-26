package internal

import (
	"fmt"
)

func ToIntList(v interface{}) ([]int, error) {
	switch v := v.(type) {
	case []interface{}:
		resp := []int{}
		for _, vv := range v {
			vvv, err := ToInt(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	case []int:
		return v, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []int", v, v)
	}
}
