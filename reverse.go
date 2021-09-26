package lambda

func (r *Object) Reverse() *Object {
	if r.err != nil {
		return r
	}

	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		r.err = err
		return r
	}

	res := []interface{}{}
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	r.obj = res
	return r
}
