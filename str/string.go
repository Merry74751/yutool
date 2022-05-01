package str

import (
	"fmt"
	"github.com/Merry74751/yutool/anyutil"
	"github.com/Merry74751/yutool/convert"
	"log"
	"strconv"
	"strings"
)

const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

func ColorText(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

func IsEmpty(str string) bool {
	return str == "" || anyutil.IsNil(str)
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		sf := fmt.Sprintf("%s convert to int error: %s", str, err)
		sf = ColorText(Red, sf)
		log.Printf(sf)
	}
	return i
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func ToBytes(str string) []byte {
	return []byte(str)
}

func ToFloat(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		sf := fmt.Sprintf("%s parse to float error: %s", s, err)
		sf = ColorText(Red, sf)
		log.Printf(sf)
	}
	return float
}

func Format(str string, params ...any) string {
	if strings.Contains(str, "{}") && len(params) != 0 {
		for _, item := range params {
			str = strings.Replace(str, "{}", convert.ToString(item), 1)
		}
	}
	return str
}

func SubString(str string, startIndex, endIndex int) string {
	lens := len(str)
	if startIndex > lens || endIndex > lens {
		return ""
	}
	if endIndex == -1 {
		endIndex = lens
	}
	bytes := str[startIndex:endIndex]
	return bytes
}

// ConvertUnderline 将字符串转为下划线形式.
// eg: Input HelloWorld return hello_world
func ConvertUnderline(str string) string {
	if IsEmpty(str) {
		return ""
	}

	bytes := []byte(str)
	builder := strings.Builder{}

	head := bytes[0]
	if head >= 65 && head <= 90 {
		head = head + 32
	}
	builder.WriteByte(head)

	for i := 1; i < len(bytes); i++ {
		b := bytes[i]
		if 65 <= b && b <= 90 {
			builder.WriteByte(95)
			b = b + 32
		}
		builder.WriteByte(b)
	}

	return builder.String()
}
