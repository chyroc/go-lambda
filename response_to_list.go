package lambda

import (
	"fmt"

	"github.com/chyroc/go-lambda/internal"
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
		v, err := internal.ToInt(item)
		if err != nil {
			return err
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
		resp, err := internal.ToIntList(item)
		if err != nil {
			return err
		}
		res = append(res, resp)
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
		k, err := internal.ToInt(key)
		if err != nil {
			return err
		}
		v, err := internal.ToFloat64List(val)
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
