package utils

// VerifyUint 校验 uint
func VerifyUint(val int, def int) int {
	if val <= 0 {
		val = def
	}
	return val
}
