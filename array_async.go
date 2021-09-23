package lambda

func (r *Object) ArrayAsync(f func(idx int, v interface{}) interface{}) *Object {
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
			objs[i] = f(i, v)
		}(i, v)
	}
	r.wg.Wait()

	r.obj = objs
	return r
}
