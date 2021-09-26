package lambda

func (r *Object) WithErr(err error) *Object {
	return r.clone(nil, err)
}
