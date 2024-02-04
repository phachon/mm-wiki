package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// Convert 转换函数
var Convert = NewConvert()

type convert struct{}

// NewConvert 创建转换函数
func NewConvert() *convert {
	return &convert{}
}

// BoolToString bool 转化为字符串
func (c *convert) BoolToString(boolValue bool) string {
	if boolValue {
		return "true"
	} else {
		return "false"
	}
}

// BoolToInt bool 转化为 int
func (c *convert) BoolToInt(boolValue bool) int {
	if boolValue {
		return 1
	} else {
		return 0
	}
}

// IntToBool int 转化为 bool
func (c *convert) IntToBool(number int) bool {
	if number == 0 {
		return false
	} else {
		return true
	}
}

// IntToString int 转化为字符串 base 范围 2-32 进制
func (c *convert) IntToString(number int64, base int) string {
	return strconv.FormatInt(number, base)
}

// StringToInt string to int(10进制)
func (c *convert) StringToInt(str string) int {
	intValue, _ := strconv.Atoi(str)
	return intValue
}

// StringToInt64 string to int64(10进制)
func (c *convert) StringToInt64(str string) int64 {
	intValue, _ := strconv.ParseInt(str, 10, 64)
	return intValue
}

// IntToString int 转化为10进制字符串 IntToString(number, 10)
func (c *convert) IntToTenString(number int) string {
	return strconv.Itoa(number)
}

// FloatToString float 转化为字符串
func (c *convert) FloatToString(f float64, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(f, fmt, prec, bitSize)
}

// ToInt64 转化任何的数为 int64
func (c *convert) ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// StringToFloat32 字符串转换为 float32
func (c *convert) StringToFloat32(str string) float32 {
	value, _ := strconv.ParseFloat(str, 32)
	return float32(value)
}

// StringToFloat64 字符串转换为 float64
func (c *convert) StringToFloat64(str string) float64 {
	value, _ := strconv.ParseFloat(str, 64)
	return value
}

// StringsToInt64 string 列表转 int64 列表
func (c *convert) StringsToInt64(strs []string) []int64 {
	if len(strs) == 0 {
		return []int64{}
	}
	var resInt64 []int64
	for _, str := range strs {
		strInt64 := c.StringToInt64(str)
		resInt64 = append(resInt64, strInt64)
	}
	return resInt64
}

// StringsToMap string 列表转 map
func (c *convert) StringsToMap(strs []string) map[string]struct{} {
	if len(strs) == 0 {
		return make(map[string]struct{})
	}
	var res = make(map[string]struct{})
	for _, str := range strs {
		res[str] = struct{}{}
	}
	return res
}
