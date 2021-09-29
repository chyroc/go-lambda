package lambda

import (
	"fmt"

	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) ToInterfaceSlice() (res []interface{}, err error) {
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

func (r *Object) ToIntListSlice() (res [][]int, err error) {
	transfer := func(idx int, item interface{}) error {
		resp, err := internal.ToIntSlice(item)
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

func (r *Object) ToMapInt2Float64Slice() (res map[int][]float64, err error) {
	res = map[int][]float64{}
	transfer := func(key, val interface{}) error {
		k, err := internal.ToInt(key)
		if err != nil {
			return err
		}
		v, err := internal.ToFloat64Slice(val)
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
