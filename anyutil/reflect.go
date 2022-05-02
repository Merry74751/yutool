package anyutil

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

// Field 根据结构体的字段名获取结构体字段的Value，如果v不是结构体返回空Value
func Field(v any, fieldName string) reflect.Value {
	if IsStruct(v) {
		return Value(v).FieldByName(fieldName)
	}
	log.Printf("%v is not struct", v)
	return reflect.Value{}
}

// FieldByIndex 获取指定值的指定索引的Field
func FieldByIndex(v any, index int) reflect.Value {
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

// Fields 获取结构体的所有Field
func Fields(v any) []reflect.Value {
	if IsStruct(v) {
		l := StructNumField(v)
		values := make([]reflect.Value, l)
		for i := 0; i < l; i++ {
			values[i] = FieldByIndex(v, i)
		}
		return values
	}
	return nil
}

// StructFieldKinds 获取结构体的所有Field的Kind
func StructFieldKinds(v any) []reflect.Kind {
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

// StructNumField 获取结构体的字段数量
func StructNumField(v any) int {
	if IsStruct(v) {
		return Value(v).NumField()
	}
	return -1
}

// FieldByName 根据结构体的字段名获取Field
func FieldByName(v any, field string) reflect.Value {
	if IsStruct(v) {
		value := Value(v)
		return value.FieldByName(field)
	}
	return reflect.Value{}
}

// StructFieldNames 获取结构体的所有字段名
func StructFieldNames(v any) []string {
	if IsStruct(v) {
		t := Type(v)
		l := StructNumField(v)
		names := make([]string, l)
		for i := 0; i < l; i++ {
			names[i] = t.Field(i).Name
		}
		return names
	}
	return nil
}

// StructFieldTag 获取结构的指定字段的标签
func StructFieldTag(v any, field, key string) string {
	if IsStruct(v) {
		if fields, b := reflect.TypeOf(v).FieldByName(field); b {
			fields.Tag.Get(key)
		}
	}
	return ""
}

// SliceType 获取切片的类型
func SliceType(v any) reflect.Type {
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

func StructField(v any, index int) reflect.StructField {
	types := Type(v)
	field := types.Field(index)
	return field
}

func StructFieldByName(v any, name string) reflect.StructField {
	types := Type(v)
	field, exits := types.FieldByName(name)
	if exits {
		return field
	}
	return reflect.StructField{}
}

func Methods(v any) []reflect.Value {
	value := Value(v)
	numMethod := value.NumMethod()
	methods := make([]reflect.Value, numMethod)
	for i := 0; i < numMethod; i++ {
		method := value.Method(i)
		methods[i] = method
	}
	return methods
}

func MethodByName(v any, methodName string) reflect.Value {
	value := Value(v)
	method := value.MethodByName(methodName)
	if !method.IsValid() {
		log.Panicf("%T dose not have method: %s", v, method)
	}
	return method
}

func InvokeMethod(v any, methodName string, param ...any) []any {
	method := MethodByName(v, methodName)
	length := len(param)

	if length == 0 {
		result := method.Call(nil)
		return convertResult(result)
	}

	values := make([]reflect.Value, length)
	for index, item := range param {
		values[index] = reflect.ValueOf(item)
	}

	result := method.Call(values)
	return convertResult(result)
}

func convertResult(result []reflect.Value) []any {
	length := len(result)
	if length == 0 {
		return []any{}
	}
	values := make([]any, length)
	for index, item := range result {
		values[index] = item.Interface()
	}
	return values
}
