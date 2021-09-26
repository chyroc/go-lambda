package lambda

// Compact Creates an array with all falsey values removed.
// The values false, null, 0, "", undefined, and NaN are falsey.
func (r *Object) Compact() *Object {
	objs := []interface{}{}
	transfer := func(idx int, item interface{}) error {
		if !isFalsey(item) {
			objs = append(objs, item)
		}
		return nil
	}

	if err := r.eachArray(transfer); err != nil {
		r.err = err
		return r
	}

	r.obj = objs
	return r
}

// The values false, null, 0, "", undefined, and NaN are falsey.
func isFalsey(v interface{}) bool {
	switch v := v.(type) {
	case error:
		return v == nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return v == 0
	case float32, float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return !v
	case interface{}:
		return v == nil
	}
	return v == nil
}
