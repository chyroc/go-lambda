package lambda

func (r *Object) FilterArray(f func(idx int, obj interface{}) bool) *Object {
	objs := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		if f(idx, obj) {
			objs = append(objs, obj)
		}
		return nil
	}

	err := r.eachArray(transfer)
	return r.clone(objs, err)
}
