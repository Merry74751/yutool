package httpstatus

import (
	"encoding/json"
	"log"
)

const (
	OK                = 200
	BadRequest        = 400
	UnAuthorized      = 401
	PermanentRequired = 402
	Forbidden         = 403
	NotFound          = 404
	ServerError       = 500
	BadGateway        = 502
)

type Res struct {
	Status  int    `json:"status,omitempty"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r Res) String() string {
	bytes, err := json.Marshal(r)
	if err != nil {
		log.Printf("json marshal error: %s", err)
	}
	return string(bytes)
}

func Success(data any) string {
	r := Res{Status: OK, Message: "请求成功", Data: data}
	return r.String()
}

func Error(status int, str string) string {
	r := Res{Status: status, Message: str}
	return r.String()
}

func Result(status int, message string) string {
	r := Res{Status: status, Message: message}
	return r.String()
}

type ListRes struct {
	Status  int
	List    any
	Message string
	total   int
}

func (l ListRes) String() string {
	bytes, err := json.Marshal(l)
	if err != nil {
		log.Printf("json marshal error: %s", err)
	}
	return string(bytes)
}

func ListResult(list any, total int) string {
	l := ListRes{List: list, Status: OK, total: total, Message: "请求成功"}
	return l.String()
}
