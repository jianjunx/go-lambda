package utils

import "strconv"

func StrToInt(str string, defValue int) int {
	if str == "" {
		return defValue
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return defValue
	}
	return val
}
