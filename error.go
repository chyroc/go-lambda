package lambda

// WithErr Set error to the object
func (r *Object) WithErr(err error) *Object {
	return r.clone(nil, err)
}
