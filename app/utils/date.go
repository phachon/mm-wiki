package utils

import (
	"github.com/astaxie/beego"
	"time"
)

var Date = NewDate()

type date struct{}

func NewDate() *date {
	return &date{}
}

//格式化 unix 时间戳
func (date *date) Format(unixTime interface{}, format string) string {
	convert := NewConvert()
	var convertTime int64
	switch unixTime.(type) {
	case string:
		convertTime = convert.StringToInt64(unixTime.(string))
	case int:
		convertTime = int64(unixTime.(int))
	case int8:
		convertTime = int64(unixTime.(int8))
	case int16:
		convertTime = int64(unixTime.(int16))
	case int32:
		convertTime = int64(unixTime.(int32))
	}
	return beego.Date(time.Unix(convertTime, 0), format)
}
