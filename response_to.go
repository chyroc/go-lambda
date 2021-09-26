package lambda

import (
	"fmt"
)

func (r *Object) ToInt() (res int, err error) {
	if r.err != nil {
		return 0, r.err
	}
	v, ok := r.obj.(int)
	if !ok {
		return 0, fmt.Errorf("%T is not int", r.obj)
	}

	return v, nil
}
