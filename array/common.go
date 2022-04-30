package array

import "github.com/Merry74751/yutool/anyutil"

func IsEmpty(arr []any) bool {
	return len(arr) == 0 || anyutil.IsNil(arr)
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
