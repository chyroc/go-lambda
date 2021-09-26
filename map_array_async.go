package lambda

func (r *Object) ArrayAsync(f func(idx int, obj interface{}) interface{}) *Object {
	return r.ArrayAsyncWithErr(func(idx int, obj interface{}) (interface{}, error) {
		return f(idx, obj), nil
	})
}

func (r *Object) ArrayAsyncWithErr(f func(idx int, obj interface{}) (interface{}, error)) *Object {
	if r.err != nil {
		return r
	}

	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		r.err = err
		return r
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

	r.obj = objs
	return r
}
