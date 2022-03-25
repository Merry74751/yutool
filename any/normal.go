package any

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
