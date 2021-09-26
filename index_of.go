package lambda

func (r *Object) IndexOf(obj interface{}) *Object {
	var objs interface{} = int(-1)
	transfer := func(idx int, item interface{}) error {
		if item == obj {
			objs = int(idx)
			return errBreak
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
