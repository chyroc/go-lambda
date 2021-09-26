package lambda

import (
	"errors"
	"fmt"
)

func (r *Object) MapArray(f func(idx int, obj interface{}) interface{}) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		objs = append(objs, f(idx, obj))
		return nil
	}

	if err := r.eachArray(transfer); err != nil {
		r.err = err
		return r
	}

	r.obj = objs
	return r
}

var errBreak = fmt.Errorf("each break")

func (r *Object) eachArray(f func(idx int, item interface{}) error) error {
	if r.err != nil {
		return r.err
	}

	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		return err
	}

	for i, v := range arr {
		if err := f(i, v); err != nil {
			if errors.Is(err, errBreak) {
				return nil
			}
			return err
		}
	}
	return nil
}

func (r *Object) eachMap(f func(key interface{}, val interface{}) error) error {
	if r.err != nil {
		return r.err
	}

	maps, err := interfaceToInterfaceMap(r.obj)
	if err != nil {
		return err
	}

	for key, val := range maps {
		if err := f(key, val); err != nil {
			if errors.Is(err, errBreak) {
				return nil
			}
			return err
		}
	}
	return nil
}
