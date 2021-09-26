package lambda

import (
	"reflect"
)

func (r *Object) Flatten() *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		vv := reflect.ValueOf(obj)
		switch vv.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < vv.Len(); j++ {
				objs = append(objs, vv.Index(j).Interface())
			}
		default:
			objs = append(objs, obj)
		}
		return nil
	}

	err := r.eachArray(transfer)
	return r.clone(objs, err)
}
