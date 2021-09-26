package lambda

import (
	"reflect"
)

type objType int

const (
	objTypeArray objType = 1
	objTypeMap   objType = 2
)

func (r *Object) objType() objType {
	vv := reflect.ValueOf(r.obj)

	switch vv.Kind() {
	case reflect.String, reflect.Array, reflect.Slice:
		return objTypeArray
	case reflect.Map:
		return objTypeMap
	default:
		return 0
	}
}
