package date

import (
	"database/sql/driver"
	"time"
)

const (
	dateTimeFormat = "2006-01-02 15:04:05"
	dateFormat     = "2006-01-02"
)

type LocalDateTime time.Time

func (t *LocalDateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.Parse(`"`+dateTimeFormat+`"`, string(data))
	if err != nil {
		now, err = time.Parse(`"`+dateFormat+`"`, string(data))
	}
	*t = LocalDateTime(now)
	return err
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalDateTime) Value() (driver.Value, error) {
	s := t.String()
	if s == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return s, nil
}

func (t *LocalDateTime) Scan(v any) error {
	now, err := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalDateTime(now)
	return err
}

func (t LocalDateTime) String() string {
	return time.Time(t).Format(dateTimeFormat)
}
