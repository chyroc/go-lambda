package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) Compact() *Object {
	objs := []interface{}{}
	transfer := func(idx int, item interface{}) error {
		if !internal.IsZero(item) {
			objs = append(objs, item)
		}
		return nil
	}

	err := r.eachArray(transfer)
	return r.clone(objs, err)
}
