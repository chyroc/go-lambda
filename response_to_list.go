package lambda

import (
	"fmt"
)

func (r *Object) ToStringList() (res []string, err error) {
	transfer := func(idx int, item interface{}) error {
		res = append(res, fmt.Sprintf("%s", item))
		return nil
	}

	if err = r.eachArray(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToIntList() (res []int, err error) {
	transfer := func(idx int, item interface{}) error {
		v, ok := item.(int)
		if !ok {
			return fmt.Errorf("%T is not int", item)
		}
		res = append(res, v)
		return nil
	}

	if err = r.eachArray(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToBoolList() (res []bool, err error) {
	transfer := func(idx int, item interface{}) error {
		v, ok := item.(bool)
		if !ok {
			return fmt.Errorf("%T is not bool", item)
		}
		res = append(res, v)
		return nil
	}

	if err = r.eachArray(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToInterfaceList() (res []interface{}, err error) {
	transfer := func(idx int, item interface{}) error {
		v, ok := item.(interface{})
		if !ok {
			return fmt.Errorf("%T is not interface{}", item)
		}
		res = append(res, v)
		return nil
	}

	if err = r.eachArray(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToIntListList() (res [][]int, err error) {
	transfer := func(idx int, item interface{}) error {
		switch obj := item.(type) {
		case []interface{}:
			item := []int{}
			for _, objChild := range obj {
				v, ok := objChild.(int)
				if !ok {
					return fmt.Errorf("%T is not int", obj)
				}
				item = append(item, v)
			}
			res = append(res, item)
		case []int:
			res = append(res, obj)
		default:
			return fmt.Errorf("%T unsupport ToIntListList", obj)
		}
		return nil
	}

	if err = r.eachArray(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToMapInt2Float64List() (res map[int][]float64, err error) {
	res = map[int][]float64{}
	transfer := func(key, val interface{}) error {
		k, ok := key.(int)
		if !ok {
			return fmt.Errorf("key %T is not int", key)
		}
		v, err := toFloat64List(val)
		if err != nil {
			return err
		}
		res[k] = v
		return nil
	}

	if err = r.eachMap(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func toFloat64List(v interface{}) (res []float64, err error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := toFloat(vv)
			if err != nil {
				return nil, err
			}
			res = append(res, vvv)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("%v(%T) is not float64 list", v, v)
	}
}

func toFloat(v interface{}) (float64, error) {
	switch v := v.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("%v(%T) is not float64", v, v)
	}
}
