package utils

import "strconv"

func StringToInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	return v, err
}

func StringToInt32(s string) (int32, error) {
	v, err := strconv.Atoi(s)
	return int32(v), err
}
