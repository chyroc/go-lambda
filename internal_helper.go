package lambda

import (
	"fmt"
	"reflect"
	"sync"
)

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
