package lambda

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) ToError() error {
	return r.err
}

func (r *Object) ToObj() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.obj, nil
}

func (r *Object) ToJoin(sep string) (string, error) {
	arr, err := r.ToStringSlice()
	if err != nil {
		return "", err
	}
	return strings.Join(arr, sep), nil
}

func (r *Object) ToObject(resp interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	arr, err := internal.ToInterfaceList(r.obj)
	if err != nil {
		return err
	}
	respV := reflect.ValueOf(resp)
	respT := reflect.TypeOf(resp)
	if respV.Kind() != reflect.Ptr {
		return fmt.Errorf("resp must be ptr")
	}
	respV = respV.Elem()
	respT = respT.Elem()
	if respV.Kind() != reflect.Slice {
		return fmt.Errorf("resp must be ptr of slice")
	}
	for i := 0; i < len(arr); i++ {
		objV := reflect.NewAt(respT.Elem().Elem(), unsafe.Pointer(reflect.ValueOf(arr[i]).Elem().UnsafeAddr()))
		respV.Set(reflect.Append(respV, objV))
	}
	return nil
}

func (r *Object) ToInterfaceSlice() (res []interface{}, err error) {
	transfer := func(idx int, item interface{}) error {
		v, ok := item.(interface{})
		if !ok {
			return fmt.Errorf("%T is not interface{}", item)
		}
		res = append(res, v)
		return nil
	}

	if err = r.eachList(transfer); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Object) ToExpectType(expectVal interface{}) (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}

	return internal.ToExpectTypeInterface(r.obj, expectVal)
}
