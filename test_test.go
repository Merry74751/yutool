package yutool

import (
	str2 "github.com/Merry74751/yutool/str"
	"testing"
)

func TestName(t *testing.T) {
	str := "abcdef"
	sub := str2.Sub(str, 2, -1)
	t.Log(sub)
}
