package anyutil

import "reflect"

func IsNil(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	if kind == reflect.Pointer {
		return reflect.ValueOf(v).Elem().IsZero()
	} else {
		return reflect.ValueOf(v).IsZero()
	}
}

func IsNotNil(v any) bool {
	return !IsNil(v)
}

func IsStruct(v any) bool {
	if Kind(v) == reflect.Struct {
		return true
	} else if Kind(v) == reflect.Ptr {
		kind := reflect.TypeOf(v).Elem().Kind()
		if kind == reflect.Struct {
			return true
		}
	}
	return false
}

func IsSlice(v any) bool {
	kind := RealKind(v)
	if kind == reflect.Slice {
		return true
	}
	return false
}

func IsArray(v any) bool {
	kind := RealKind(v)
	if kind == reflect.Array {
		return true
	}
	return false
}
