package lambda

import (
	"reflect"
)

func (r *Object) Flatten() *Object {
	if r.err != nil {
		return r
	}

	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		r.err = err
		return r
	}

	objs := []interface{}{}
	for _, v := range arr {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < vv.Len(); j++ {
				objs = append(objs, vv.Index(j).Interface())
			}
		default:
			objs = append(objs, v)
		}
	}
	r.obj = objs
	return r
}
