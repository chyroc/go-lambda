package lambda

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func (r *Object) Error() error {
	return r.err
}

func (r *Object) Obj() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.obj, nil
}

func (r *Object) Join(sep string) (string, error) {
	arr, err := r.StringList()
	if err != nil {
		return "", err
	}
	return strings.Join(arr, sep), nil
}

func (r *Object) String() (res string, err error) {
	if r.err != nil {
		return "", r.err
	}
	switch v := r.obj.(type) {
	case []rune, []byte, string:
		return toStrictString(r.obj)
	case []interface{}:
		if vv, err := interfaceList2Int32List(v); err == nil {
			return string(vv), nil
		}
		if vv, err := interfaceList2Uint8List(v); err == nil {
			return string(vv), nil
		}
		return "", fmt.Errorf("list of %T unsupport .String lambda operator", v[0])
	default:
		return "", fmt.Errorf("%T unsupport .String lambda operator", r.obj)
	}
}

func (r *Object) StringList() (res []string, err error) {
	if r.err != nil {
		return nil, r.err
	}
	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		return nil, err
	}
	for _, v := range arr {
		res = append(res, fmt.Sprintf("%s", v))
	}
	return res, nil
}

func (r *Object) ToList(resp interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	arr, err := interfaceToInterfaceList(r.obj)
	if err != nil {
		return err
	}
	respV := reflect.ValueOf(resp)
	respT := reflect.TypeOf(resp)
	if respV.Kind() != reflect.Ptr {
		return fmt.Errorf("resp must be ptr")
	}
	respV = respV.Elem()
	respT = respT.Elem()
	if respV.Kind() != reflect.Slice {
		return fmt.Errorf("resp must be ptr of slice")
	}
	for i := 0; i < len(arr); i++ {
		objV := reflect.NewAt(respT.Elem().Elem(), unsafe.Pointer(reflect.ValueOf(arr[i]).Elem().UnsafeAddr()))
		respV.Set(reflect.Append(respV, objV))
	}
	return nil
}
