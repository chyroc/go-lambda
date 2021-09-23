package lambda

func (r *Object) Array(f func(idx int, v interface{}) interface{}) *Object {
	if r.err != nil {
		return r
	}

	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		r.err = err
		return r
	}

	objs := []interface{}{}
	for i, v := range arr {
		res := f(i, v)
		objs = append(objs, res)
	}
	r.obj = objs
	return r
}
