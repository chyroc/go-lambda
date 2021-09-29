package lambda

func (r *Object) GroupByArray(f func(idx int, obj interface{}) interface{}) *Object {
	objs := map[interface{}][]interface{}{}
	transfer := func(idx int, item interface{}) error {
		key := f(idx, item)
		objs[key] = append(objs[key], item)
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}
