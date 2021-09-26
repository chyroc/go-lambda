package lambda

func (r *Object) Compact() *Object {
	objs := []interface{}{}
	transfer := func(idx int, item interface{}) error {
		if !isFalsey(item) {
			objs = append(objs, item)
		}
		return nil
	}

	err := r.eachArray(transfer)
	return r.clone(objs, err)
}
