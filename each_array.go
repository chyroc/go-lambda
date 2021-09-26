package lambda

func (r *Object) EachArray(f func(idx int, obj interface{})) error {
	return r.eachArray(func(idx int, item interface{}) error {
		f(idx, item)
		return nil
	})
}
