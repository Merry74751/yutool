package any

import (
	"log"
	"reflect"
)

// Kind 获取指定值的类型的分类
func Kind(v any) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

// RealKind 获取指针的类型的分类
func RealKind(v any) reflect.Kind {
	kind := Kind(v)
	if kind == reflect.Ptr {
		return reflect.TypeOf(v).Elem().Kind()
	}
	return kind
}

// Type 获取指定值的类型
func Type(v any) reflect.Type {
	if Kind(v) == reflect.Ptr {
		return reflect.TypeOf(v).Elem()
	}
	return reflect.TypeOf(v)
}

// Value 获取指定值的反射值
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

// GetFieldByIndex 获取指定值的指定索引的Field
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

// IsStruct 判断指定值是否是结构体类型
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

// Fields 获取结构体的所有Field
func Fields(v any) []reflect.Value {
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

// GetStructFieldKinds 获取结构体的所有Field的Kind
func GetStructFieldKinds(v any) []reflect.Kind {
	fields := Fields(v)
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

// GetStructNumField 获取结构体的字段数量
func GetStructNumField(v any) int {
	if IsStruct(v) {
		return Value(v).NumField()
	}
	return -1
}

// GetStructFieldByName 根据结构体的字段名获取Field
func GetStructFieldByName(v any, field string) reflect.Value {
	if IsStruct(v) {
		value := Value(v)
		return value.FieldByName(field)
	}
	return reflect.Value{}
}

// GetStructFieldNames 获取结构体的所有字段名
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

// GetStructFieldTag 获取结构的指定字段的标签
func GetStructFieldTag(v any, field, key string) string {
	if IsStruct(v) {
		if fields, b := reflect.TypeOf(v).FieldByName(field); b {
			fields.Tag.Get(key)
		}
	}
	return ""
}

// GetSliceType 获取切片的类型
func GetSliceType(v any) reflect.Type {
	if RealKind(v) == reflect.Slice {
		value := Value(v)
		return value.Type().Elem()
	}
	return nil
}

func StructFields(v any) []reflect.StructField {
	typ := Type(v)
	numField := typ.NumField()
	fields := make([]reflect.StructField, numField)
	for i := 0; i < numField; i++ {
		structField := typ.Field(i)
		fields[i] = structField
	}
	return fields
}
