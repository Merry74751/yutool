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

func Sub(str string, startIndex, endIndex int) string {
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
