package lambda

func (r *Object) Transfer(f func(obj interface{}) interface{}) *Object {
	if r.err != nil {
		return r.clone(nil, r.err)
	}

	return r.clone(f(r.obj), nil)
}
