package lambda

func (r *Object) GroupByArray(f func(idx int, obj interface{}) interface{}) *Object {
	objs := map[interface{}][]interface{}{}
	transfer := func(idx int, item interface{}) error {
		key := f(idx, item)
		objs[key] = append(objs[key], item)
		return nil
	}

	if err := r.eachArray(transfer); err != nil {
		r.err = err
		return r
	}
	r.obj = objs
	return r
}
