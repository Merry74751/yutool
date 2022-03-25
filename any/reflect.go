package any

import (
	"log"
	"reflect"
)

func Kind(v any) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

func Value(v any) reflect.Value {
	if Kind(v) == reflect.Ptr {
		return reflect.ValueOf(v).Elem()
	}
	return reflect.ValueOf(v)
}

func GetField(v any, fieldName string) reflect.Value {
	if !IsStruct(v) {
		log.Printf("%v is not struct", v)
		return reflect.Value{}
	}
	return Value(v).FieldByName(fieldName)
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
