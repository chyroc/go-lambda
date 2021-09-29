package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) ToIntArray() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToIntArray(r.obj)
}

func (r *Object) ToInt8Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt8Array(r.obj)
}

func (r *Object) ToInt16Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt16Array(r.obj)
}

func (r *Object) ToInt32Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt32Array(r.obj)
}

func (r *Object) ToInt64Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt64Array(r.obj)
}

func (r *Object) ToUintArray() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUintArray(r.obj)
}

func (r *Object) ToUint8Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint8Array(r.obj)
}

func (r *Object) ToUint16Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint16Array(r.obj)
}

func (r *Object) ToUint32Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint32Array(r.obj)
}

func (r *Object) ToUint64Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint64Array(r.obj)
}

func (r *Object) ToUintptrArray() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUintptrArray(r.obj)
}

func (r *Object) ToFloat32Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToFloat32Array(r.obj)
}

func (r *Object) ToFloat64Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToFloat64Array(r.obj)
}

func (r *Object) ToBoolArray() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToBoolArray(r.obj)
}

func (r *Object) ToComplex64Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToComplex64Array(r.obj)
}

func (r *Object) ToComplex128Array() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToComplex128Array(r.obj)
}

func (r *Object) ToStringArray() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToStringArray(r.obj)
}
