package yutool

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	u := User{}
	t.Log(reflect.TypeOf(u).Kind())
}

type User struct {
}
