package lambda

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func (r *Obj) Error() error {
	return r.err
}

func (r *Obj) Obj() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.obj, nil
}

func (r *Obj) Join(sep string) (string, error) {
	arr, err := r.StringList()
	if err != nil {
		return "", err
	}
	return strings.Join(arr, sep), nil
}

func (r *Obj) String() (res string, err error) {
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

func (r *Obj) StringList() (res []string, err error) {
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

func (r *Obj) ToList(resp interface{}) (err error) {
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

func toStrictString(v interface{}) (string, error) {
	switch v := v.(type) {
	case []rune:
		return string(v), nil
	case []byte:
		return string(v), nil
	case string:
		return string(v), nil
	// case []interface{}:
	// 	for _, vv := range v {
	//
	// 	}
	default:
		return "", fmt.Errorf("%T cannot convert to string", v)
	}
}

func interfaceList2Int32List(vv []interface{}) (res []int32, err error) {
	for _, v := range vv {
		if vvv, ok := v.(int32); ok {
			res = append(res, vvv)
		} else {
			return nil, fmt.Errorf("%T cannot convert to int32 of []int32", v)
		}
	}
	return res, nil
}

func interfaceList2Uint8List(vv []interface{}) (res []uint8, err error) {
	for _, v := range vv {
		if vvv, ok := v.(uint8); ok {
			res = append(res, vvv)
		} else {
			return nil, fmt.Errorf("%T cannot convert to uint8 of []uint8", v)
		}
	}
	return res, nil
}
