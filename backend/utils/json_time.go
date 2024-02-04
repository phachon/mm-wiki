package utils

import (
	"fmt"
	"time"

	"database/sql/driver"
)

// JsonTime json 格式化的时间
type JsonTime struct {
	time.Time
}

// NewJsonTime 创建一个 json time
func NewJsonTime(time time.Time) JsonTime {
	return JsonTime{
		time,
	}
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (jt JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", jt.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (jt JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jt.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jt.Time, nil
}

// Scan valueof time.Time
func (jt *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jt = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
