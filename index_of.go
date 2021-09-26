package lambda

func (r *Object) IndexOf(obj interface{}) *Object {
	var objs interface{} = int(-1)
	transfer := func(idx int, item interface{}) error {
		if item == obj {
			objs = int(idx)
			return ErrBreak
		}
		return nil
	}

	err := r.eachArray(transfer)
	return r.clone(objs, err)
}
