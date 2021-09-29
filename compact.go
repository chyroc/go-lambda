package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

// Compact Remove 0-valued elements from the list
//
// Compact [0, 1, false, true, "", "str"], will return [1, true, "str"]
func (r *Object) Compact() *Object {
	objs := []interface{}{}
	transfer := func(idx int, item interface{}) error {
		if !internal.IsZero(item) {
			objs = append(objs, item)
		}
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}
