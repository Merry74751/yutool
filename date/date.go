package date

import (
	"log"
	"time"
)

const (
	Second = time.Second
	Minute = time.Minute
	Hour   = time.Hour
	Day    = Hour * 24
	Week   = Day * 7
)

func Now() time.Time {
	return time.Now()
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Formatter(format string) string {
	return Now().Format(format)
}

func GetYear() string {
	return Now().Format("2006")
}

func GetMonth() string {
	return Now().Format("01")
}

func GetDayOfMonth() string {
	return Now().Format("02")
}

func GetDay() string {
	return Now().Format("2006-01-02")
}

func Parse(format, value string) time.Time {
	t, err := time.Parse(format, value)
	if err != nil {
		log.Printf("parse time: %s error: %s", value, err)
	}
	return t
}

func NextWeek(t time.Time) time.Time {
	return t.Add(time.Hour * 24 * 7)
}

func NextDay(t time.Time) time.Time {
	return t.Add(time.Hour * 24)
}

func NextHour(t time.Time) time.Time {
	return t.Add(time.Hour)
}

func NextMinute(t time.Time) time.Time {
	return t.Add(time.Minute)
}

func NextMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0)
}
