package string

import (
	Any "github.com/Merry74751/yutool/any"
	"log"
	"strconv"
)

func StrIsEmpty(str string) bool {
	return str == "" || Any.IsNil(str)
}

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("%s convert to int error: %s", str, err)
	}
	return i
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}
