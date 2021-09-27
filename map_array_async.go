package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) MapArrayAsync(f func(idx int, obj interface{}) interface{}) *Object {
	return r.MapArrayAsyncWithErr(func(idx int, obj interface{}) (interface{}, error) {
		return f(idx, obj), nil
	})
}

func (r *Object) MapArrayAsyncWithErr(f func(idx int, obj interface{}) (interface{}, error)) *Object {
	if r.err != nil {
		return r
	}

	arr, err := internal.ToInterfaceList(r.obj)
	if err != nil {
		return r.clone(nil, err)
	}

	objs := make([]interface{}, len(arr))
	for i, v := range arr {
		r.wg.Add(1)
		go func(i int, v interface{}) {
			defer r.wg.Done()
			if r.err != nil {
				return
			}
			objs[i], err = f(i, v)
			if err != nil {
				r.err = err
				return
			}
		}(i, v)
	}
	r.wg.Wait()

	return r.clone(objs, nil)
}
