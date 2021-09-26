package lambda

func (r *Object) Chunk(size int) *Object {
	objs := []interface{}{}
	item := []interface{}{}
	transfer := func(idx int, obj interface{}) error {
		if len(item) < size {
			item = append(item, obj)
		} else {
			objs = append(objs, item)
			item = []interface{}{obj}
		}
		return nil
	}

	if err := r.eachArray(transfer); err != nil {
		return r.clone(nil, err)
	}

	if len(item) > 0 {
		objs = append(objs, item)
		item = []interface{}{}
	}

	return r.clone(objs, nil)
}
