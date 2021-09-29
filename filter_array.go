package lambda

// FilterList Traverse each element in the array or slice, use the filter callback function to call this element,
// the new object will return all the elements that make the filter function return true.
//
// Use "is even" to call the filter function, then [1, 2, 3, 4] will return [2 4]
func (r *Object) FilterList(f func(idx int, obj interface{}) bool) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		if f(idx, obj) {
			objs = append(objs, obj)
		}
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}
