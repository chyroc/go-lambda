package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

func (r *Object) ToInt() (int, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToInt(r.obj)
}

func (r *Object) ToInt8() (int8, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToInt8(r.obj)
}

func (r *Object) ToInt16() (int16, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToInt16(r.obj)
}

func (r *Object) ToInt32() (int32, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToInt32(r.obj)
}

func (r *Object) ToInt64() (int64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToInt64(r.obj)
}

func (r *Object) ToUint() (uint, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToUint(r.obj)
}

func (r *Object) ToUint8() (uint8, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToUint8(r.obj)
}

func (r *Object) ToUint16() (uint16, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToUint16(r.obj)
}

func (r *Object) ToUint32() (uint32, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToUint32(r.obj)
}

func (r *Object) ToUint64() (uint64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToUint64(r.obj)
}

func (r *Object) ToFloat32() (float32, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToFloat32(r.obj)
}

func (r *Object) ToFloat64() (float64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToFloat64(r.obj)
}

func (r *Object) ToBool() (bool, error) {
	if r.err != nil {
		return false, r.err
	}
	return internal.ToBool(r.obj)
}

func (r *Object) ToComplex64() (complex64, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToComplex64(r.obj)
}

func (r *Object) ToComplex128() (complex128, error) {
	if r.err != nil {
		return 0, r.err
	}
	return internal.ToComplex128(r.obj)
}

func (r *Object) ToString() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	return internal.ToString(r.obj)
}
