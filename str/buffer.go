package str

import (
	"github.com/Merry74751/yutool/convert"
	"log"
	"strings"
)

type strBuffer struct {
	strings.Builder
}

func NewBuffer() *strBuffer {
	return new(strBuffer)
}

func (buffer *strBuffer) Append(value any) *strBuffer {
	str := convert.ToString(value)
	_, err := buffer.WriteString(str)
	if err != nil {
		log.Printf("append error: %s", err)
	}
	return buffer
}

func (buffer *strBuffer) ToString() string {
	return buffer.String()
}

func (buffer *strBuffer) ToBytes() []byte {
	return []byte(buffer.String())
}
