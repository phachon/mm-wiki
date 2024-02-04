package utils

import (
	"fmt"
	"time"

	"database/sql/driver"
)

// JsonDate json 格式化的日期
type JsonDate struct {
	time.Time
}

// NewJsonDate 创建一个 json time
func NewJsonDate(time time.Time) JsonDate {
	return JsonDate{
		time,
	}
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d
func (jt JsonDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", jt.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (jt JsonDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jt.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jt.Time, nil
}

// Scan valueof time.Time
func (jt *JsonDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jt = JsonDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to date", v)
}
