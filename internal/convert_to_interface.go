package internal

import (
	"fmt"
	"reflect"
)

func ToExpectTypeInterface(v interface{}, expectTypeVal interface{}) (resp interface{}, err error) {
	et := reflect.TypeOf(expectTypeVal)
	return toExpectTypeInterface(v, et)
}

func toExpectTypeInterface(v interface{}, et reflect.Type) (resp interface{}, err error) {
	vv := reflect.ValueOf(v)
	// vt := reflect.TypeOf(v)
	// ev := reflect.ValueOf(expectTypeVal)

	switch et.Kind() {
	case reflect.Invalid:
		return nil, fmt.Errorf("%v(%T) can't convert to invalid", v, v)
	case reflect.Bool:
		resp, err := ToBool(v)
		return resp, err
	case reflect.Int:
		resp, err := ToInt(v)
		return resp, err
	case reflect.Int8:
		resp, err := ToInt8(v)
		return resp, err
	case reflect.Int16:
		resp, err := ToInt16(v)
		return resp, err
	case reflect.Int32:
		resp, err := ToInt32(v)
		return resp, err
	case reflect.Int64:
		resp, err := ToInt64(v)
		return resp, err
	case reflect.Uint:
		resp, err := ToUint(v)
		return resp, err
	case reflect.Uint8:
		resp, err := ToUint8(v)
		return resp, err
	case reflect.Uint16:
		resp, err := ToUint16(v)
		return resp, err
	case reflect.Uint32:
		resp, err := ToUint32(v)
		return resp, err
	case reflect.Uint64:
		resp, err := ToUint64(v)
		return resp, err
	case reflect.Uintptr:
		resp, err := ToUintptr(v)
		return resp, err
	case reflect.Float32:
		resp, err := ToFloat32(v)
		return resp, err
	case reflect.Float64:
		resp, err := ToFloat64(v)
		return resp, err
	case reflect.Complex64:
		resp, err := ToComplex64(v)
		return resp, err
	case reflect.Complex128:
		resp, err := ToComplex128(v)
		return resp, err
	case reflect.Array:
		vtt := reflect.New(reflect.ArrayOf(vv.Len(), et.Elem())).Elem()
		for i := 0; i < vv.Len(); i++ {
			vvChild := vv.Index(i)
			eTargetChild, err := toExpectTypeInterface(vvChild.Interface(), et.Elem())
			if err != nil {
				return nil, err
			}
			vtt.Index(i).Set(reflect.ValueOf(eTargetChild))
		}
		return vtt.Interface(), nil
	case reflect.Chan:
		return nil, fmt.Errorf("%v(%T) can't convert to chan", v, v)
	case reflect.Func:
		return nil, fmt.Errorf("%v(%T) can't convert to func", v, v)
	case reflect.Interface:
		panic("unreachable interface etype")
	case reflect.Map:
		etk := et.Key()
		etv := et.Elem()
		vrange := vv.MapRange()
		vtt := reflect.MakeMapWithSize(et, vv.Len())
		for vrange.Next() {
			key, err := toExpectTypeInterface(vrange.Key().Interface(), etk)
			if err != nil {
				return nil, err
			}
			val, err := toExpectTypeInterface(vrange.Value().Interface(), etv)
			if err != nil {
				return nil, err
			}
			vtt.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
		}
		return vtt.Interface(), nil
	case reflect.Ptr:
		val, err := toExpectTypeInterface(v, et.Elem())
		if err != nil {
			return nil, err
		}
		pv := reflect.New(et.Elem())
		pv.Elem().Set(reflect.ValueOf(val))
		return pv.Interface(), nil
	case reflect.Slice:
		vtt := reflect.MakeSlice(et, vv.Len(), vv.Len())
		for i := 0; i < vv.Len(); i++ {
			vvChild := vv.Index(i)
			eTargetChild, err := toExpectTypeInterface(vvChild.Interface(), et.Elem())
			if err != nil {
				return nil, err
			}
			vtt.Index(i).Set(reflect.ValueOf(eTargetChild))
		}
		return vtt.Interface(), nil
	case reflect.String:
		resp, err := ToString(v)
		return resp, err
	case reflect.Struct:
		if vv.Kind() == reflect.Ptr {
			vv = vv.Elem()
		}
		vtt := reflect.New(et)
		for i := 0; i < et.NumField(); i++ {
			fieldET := et.Field(i)
			vvField := vv.FieldByName(fieldET.Name)
			targetField := vtt.Elem().FieldByName(fieldET.Name)
			targetFieldVal, err := toExpectTypeInterface(vvField.Interface(), fieldET.Type)
			if err != nil {
				return nil, err
			}
			targetField.Set(reflect.ValueOf(targetFieldVal))
		}
		return vtt.Elem().Interface(), nil
	case reflect.UnsafePointer:
		return nil, fmt.Errorf("%v(%T) can't convert to UnsafePointer", v, v)
	default:
		panic("unreachable")
	}
}
