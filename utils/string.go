package utils

import "strings"

func UpperFirstWord(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func LowerFirstWord(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
