package lambda

// Chunk Creates an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements.
func (r *Object) Chunk(size int) *Object {
	objs := []interface{}{}
	// res := [][]interface{}{}
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
		r.err = err
		return r
	}

	if len(item) > 0 {
		objs = append(objs, item)
		item = []interface{}{}
	}

	r.obj = objs
	return r
}
