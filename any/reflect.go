package any

import (
	"log"
	"reflect"
)

// Kind 获取结构体类型
func Kind(v any) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

func RealKind(v any) reflect.Kind {
	kind := Kind(v)
	if kind == reflect.Ptr {
		return reflect.TypeOf(v).Elem().Kind()
	}
	return kind
}

func Type(v any) reflect.Type {
	if Kind(v) == reflect.Ptr {
		return reflect.TypeOf(v).Elem()
	}
	return reflect.TypeOf(v)
}

// Value 获取结构体Value
func Value(v any) reflect.Value {
	if Kind(v) == reflect.Ptr {
		return reflect.ValueOf(v).Elem()
	}
	return reflect.ValueOf(v)
}

// GetField 根据结构体的字段名获取结构体字段的Value，如果v不是结构体返回空Value
func GetField(v any, fieldName string) reflect.Value {
	if IsStruct(v) {
		return Value(v).FieldByName(fieldName)
	}
	log.Printf("%v is not struct", v)
	return reflect.Value{}
}

func GetFieldByIndex(v any, index int) reflect.Value {
	if IsStruct(v) {
		l := Value(v).NumField()
		if index <= l {
			return Value(v).Field(index)
		} else {
			log.Printf("%d out of the struct fields index %d", index, l)
			return reflect.Value{}
		}
	}
	log.Printf("%v is not struct", v)
	return reflect.Value{}
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

func GetStructFields(v any) []reflect.Value {
	if IsStruct(v) {
		l := GetStructNumField(v)
		values := make([]reflect.Value, l)
		for i := 0; i < l; i++ {
			values[i] = GetFieldByIndex(v, i)
		}
		return values
	}
	return nil
}

func GetStructFieldKinds(v any) []reflect.Kind {
	fields := GetStructFields(v)
	if fields != nil {
		kinds := make([]reflect.Kind, len(fields))
		for i, field := range fields {
			kinds[i] = RealKind(field)
			i++
		}
		return kinds
	}
	return nil
}

func GetStructNumField(v any) int {
	if IsStruct(v) {
		return Value(v).NumField()
	}
	return -1
}

func GetStructFieldNames(v any) []string {
	if IsStruct(v) {
		t := Type(v)
		l := GetStructNumField(v)
		names := make([]string, l)
		for i := 0; i < l; i++ {
			names[i] = t.Field(i).Name
		}
		return names
	}
	return nil
}
