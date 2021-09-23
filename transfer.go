package lambda

func (r *Object) Transfer(f func(obj interface{}) interface{}) *Object {
	if r.err != nil {
		return r
	}

	r.obj = f(r.obj)

	return r
}
