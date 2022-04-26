package array

import any2 "github.com/Merry74751/yutool/any"

func IsEmpty(arr []any) bool {
	return len(arr) == 0 || any2.IsNil(arr)
}

func DoFilter(arr []Filter) []any {
	v := make([]any, 0)
	for _, item := range arr {
		if item.filter(item) {
			v = append(v, item)
		}
	}
	return v
}

type Filter interface {
	filter(v any) bool
}
