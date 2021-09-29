package lambda

// EachList Traverse the array or slice, and use the callback function to process each element
func (r *Object) EachList(f func(idx int, obj interface{})) error {
	return r.eachList(func(idx int, item interface{}) error {
		f(idx, item)
		return nil
	})
}
