package lambda

import (
	"errors"

	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) MapList(f func(idx int, obj interface{}) interface{}) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		objs = append(objs, f(idx, obj))
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}

func (r *Object) MapListErr(f func(idx int, obj interface{}) (interface{}, error)) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		res, err := f(idx, obj)
		if err != nil {
			return err
		}
		objs = append(objs, res)
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}

func (r *Object) eachList(f func(idx int, item interface{}) error) error {
	if r.err != nil {
		return r.err
	}

	arr, err := internal.ToInterfaceList(r.obj)
	if err != nil {
		return err
	}

	for i, v := range arr {
		if err := f(i, v); err != nil {
			if errors.Is(err, ErrBreak) {
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
			if errors.Is(err, ErrBreak) {
				return nil
			}
			return err
		}
	}
	return nil
}
