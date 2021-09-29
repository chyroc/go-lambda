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
