package lambda

func (r *Object) FilterArray(f func(idx int, obj interface{}) bool) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		if f(idx, obj) {
			objs = append(objs, obj)
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
