package lambda

// Chunk Split the list into shorter length lists
//
// Chunk [1, 2, 3, 4, 5] with length 2, return [[1 2] [3 4] [5]]
// If the length of the last group is less than length, then the last group will be used as an element alone
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

	if err := r.eachList(transfer); err != nil {
		return r.clone(nil, err)
	}

	if len(item) > 0 {
		objs = append(objs, item)
		item = []interface{}{}
	}

	return r.clone(objs, nil)
}
