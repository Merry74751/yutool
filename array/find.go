package array

import (
	"github.com/Merry74751/yutool/anyutil"
	"log"
)

func Find[T any](array []T, value T) T {
	exist, i := isExist(array, value)
	if exist {
		return array[i]
	}
	return nil
}

func IsExist[T any](array []T, value T) bool {
	exist, _ := isExist(array, value)
	return exist
}

func FindIndex[T any](array []T, value T) int {
	_, i := isExist(array, value)
	return i
}

func isExist[T any](array []T, value T) (bool, int) {
	if len(array) == 0 {
		log.Printf("array: %v is nil", array)
		return false, -1
	}
	v1 := anyutil.Value(value)
	for index, item := range array {
		v2 := anyutil.Value(item)
		if v1 == v2 {
			return true, index
		}
	}
	return false, -1
}
