package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) Reverse() *Object {
	if r.err != nil {
		return r.clone(nil, r.err)
	}

	arr, err := internal.ToInterfaceList(r.obj)
	if err != nil {
		return r.clone(nil, err)
	}

	res := []interface{}{}
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	return r.clone(res, err)
}
