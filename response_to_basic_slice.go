package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) ToIntSlice() ([]int, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToIntSlice(r.obj)
}

func (r *Object) ToInt8Slice() ([]int8, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt8Slice(r.obj)
}

func (r *Object) ToInt16Slice() ([]int16, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt16Slice(r.obj)
}

func (r *Object) ToInt32Slice() ([]int32, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt32Slice(r.obj)
}

func (r *Object) ToInt64Slice() ([]int64, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToInt64Slice(r.obj)
}

func (r *Object) ToUintSlice() ([]uint, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUintSlice(r.obj)
}

func (r *Object) ToUint8Slice() ([]uint8, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint8Slice(r.obj)
}

func (r *Object) ToUint16Slice() ([]uint16, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint16Slice(r.obj)
}

func (r *Object) ToUint32Slice() ([]uint32, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint32Slice(r.obj)
}

func (r *Object) ToUint64Slice() ([]uint64, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUint64Slice(r.obj)
}

func (r *Object) ToUintptrSlice() ([]uintptr, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToUintptrSlice(r.obj)
}

func (r *Object) ToFloat32Slice() ([]float32, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToFloat32Slice(r.obj)
}

func (r *Object) ToFloat64Slice() ([]float64, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToFloat64Slice(r.obj)
}

func (r *Object) ToBoolSlice() ([]bool, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToBoolSlice(r.obj)
}

func (r *Object) ToComplex64Slice() ([]complex64, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToComplex64Slice(r.obj)
}

func (r *Object) ToComplex128Slice() ([]complex128, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToComplex128Slice(r.obj)
}

func (r *Object) ToStringSlice() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.ToStringSlice(r.obj)
}
