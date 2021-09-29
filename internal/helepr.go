package internal

import (
	"reflect"
)

var intSize = 32 << (^uint(0) >> 63) // 32 or 64

func ReflectCanInterface(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return true
	}
	return false
}
