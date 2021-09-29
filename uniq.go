package lambda

func (r *Object) Uniq() *Object {
	objs := []interface{}{}
	done := map[interface{}]bool{}
	transfer := func(idx int, item interface{}) error {
		if !done[item] {
			objs = append(objs, item)
			done[item] = true
		}
		return nil
	}

	err := r.eachList(transfer)
	return r.clone(objs, err)
}
