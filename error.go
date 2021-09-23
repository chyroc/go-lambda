package lambda

func (r *Object) SetErr(err error) *Object {
	r.err = err
	return r
}
