package lambda

import (
	"fmt"
)

func (r *Object) ToStringList() (res []string, err error) {
	transfer := func(obj interface{}) error {
		res = append(res, fmt.Sprintf("%s", obj))
		return nil
	}

	if err = r.toList(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToIntList() (res []int, err error) {
	transfer := func(obj interface{}) error {
		v, ok := obj.(int)
		if !ok {
			return fmt.Errorf("%T is not int", obj)
		}
		res = append(res, v)
		return nil
	}

	if err = r.toList(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) toList(f func(obj interface{}) error) (err error) {
	if r.err != nil {
		return r.err
	}
	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		return err
	}
	for _, v := range arr {
		err := f(v)
		if err != nil {
			return err
		}
	}
	return nil
}
