package lambda

func (r *Object) Uniq() *Object {
	objs := []interface{}{}
	done := map[interface{}]bool{}
	transfer := func(idx int, item interface{}) error {
		if !done[item] {
			objs = append(objs, item)
			done[item] = true
		}
		return nil
	}

	if err := r.eachArray(transfer); err != nil {
		r.err = err
		return r
	}

	r.obj = objs
	return r
}
