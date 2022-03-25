package https

import (
	"encoding/json"
	"github.com/Merry74751/yutool/str"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func ParseBody(r *http.Request, v any) any {
	body := r.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("read body error: %s", err)
		return v
	}

	if len(bytes) == 0 {
		return v
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		log.Printf("%v convert to %v error: %s", bytes, v, err)
		return v
	}
	return v
}

func GetParam(r *http.Request, key string) string {
	query := r.URL.Query()
	return query.Get(key)
}

func GetParamInt(r *http.Request, key string) int {
	param := GetParam(r, key)
	if str.IsEmpty(param) {
		return -1
	}
	return str.ToInt(param)
}

func PathVariable(path, key string, r *http.Request) string {
	url := r.URL.Path
	s1 := strings.Split(path, "/")
	s2 := strings.Split(url, "/")
	for index, item := range s1 {
		if isPathVariable(item) {
			if strings.Contains(item, key) {
				return s2[index]
			}
		}
	}
	return ""
}

func isPathVariable(str string) bool {
	compile, err := regexp.Compile("^\\{.*\\}$")
	if err != nil {
		log.Printf("regexp error: %s", err)
	}
	return compile.MatchString(str)
}
