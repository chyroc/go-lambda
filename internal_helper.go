package lambda

import (
	"fmt"
	"reflect"
	"sync"
)

func interfaceToInterfaceList(v interface{}) (res []interface{}, err error) {
	vv := reflect.ValueOf(v)
	canToArrKinds := []reflect.Kind{
		reflect.String,
		reflect.Slice,
		reflect.Array,
	}
	if !isInKind(vv.Kind(), canToArrKinds) {
		return nil, fmt.Errorf("%T unsupport to array lambda operator", v)
	}
	switch vv.Kind() {
	case reflect.String:
		for _, v := range []rune(vv.String()) {
			res = append(res, rune(v))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < vv.Len(); i++ {
			res = append(res, vv.Index(i).Interface())
		}
	}
	return
}

func interfaceToInterfaceMap(v interface{}) (res map[interface{}]interface{}, err error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return nil, fmt.Errorf("%T unsupport to map lambda operator", v)
	}

	m := vv.MapRange()
	res = map[interface{}]interface{}{}
	for m.Next() {
		res[m.Key().Interface()] = m.Value().Interface()
	}
	return
}

func isInKind(kind reflect.Kind, kinds []reflect.Kind) bool {
	for _, v := range kinds {
		if kind == v {
			return true
		}
	}
	return false
}

func toStrictString(v interface{}) (string, error) {
	switch v := v.(type) {
	case []rune:
		return string(v), nil
	case []byte:
		return string(v), nil
	case string:
		return string(v), nil
	default:
		return "", fmt.Errorf("%T cannot convert to string", v)
	}
}

func interfaceList2Int32List(vv []interface{}) (res []int32, err error) {
	for _, v := range vv {
		if vvv, ok := v.(int32); ok {
			res = append(res, vvv)
		} else {
			return nil, fmt.Errorf("%T cannot convert to int32 of []int32", v)
		}
	}
	return res, nil
}

func interfaceList2Uint8List(vv []interface{}) (res []uint8, err error) {
	for _, v := range vv {
		if vvv, ok := v.(uint8); ok {
			res = append(res, vvv)
		} else {
			return nil, fmt.Errorf("%T cannot convert to uint8 of []uint8", v)
		}
	}
	return res, nil
}

// The values false, null, 0, "", undefined, and NaN are falsey.
func isFalsey(v interface{}) bool {
	switch v := v.(type) {
	case error:
		return v == nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return v == 0
	case float32, float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return !v
	case interface{}:
		return v == nil
	}
	return v == nil
}

func (r *Object) clone(obj interface{}, err error) *Object {
	if err != nil {
		return &Object{
			err: err,
			obj: r.obj,
			wg:  new(sync.WaitGroup),
		}
	}
	if r.err != nil {
		return &Object{
			err: r.err,
			obj: r.obj,
			wg:  new(sync.WaitGroup),
		}
	}
	return &Object{
		err: nil,
		obj: obj,
		wg:  new(sync.WaitGroup),
	}
}
