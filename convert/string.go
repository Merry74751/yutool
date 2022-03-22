package convert

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func ToString(value any) string {
	switch value.(type) {
	case string:
		return value.(string)
	default:
		kind := reflect.TypeOf(value).Kind()
		if kind == reflect.Struct {
			bytes, err := json.Marshal(value)
			if err != nil {
				log.Printf("json marshal error: %s", bytes)
				return ""
			}
			return string(bytes)
		}
		return fmt.Sprint(value)
	}
}
