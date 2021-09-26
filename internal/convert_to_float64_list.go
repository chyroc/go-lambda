package internal

import (
	"fmt"
)

func ToFloat64List(v interface{}) (resp []float64, err error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := ToFloat(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("%v(%T) can't convert to []float64", v, v)
	}
}
