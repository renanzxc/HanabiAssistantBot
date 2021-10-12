package main

import (
	"log"
	"strconv"
	"strings"
)

func parseTelegramData(data string) []string {
	return strings.Split(data, "|")
}

func ValueToString(value interface{}) string {
	switch value := value.(type) {
	case *int:
		return strconv.FormatInt(int64(*value), 10)
	case *ColorType:
		return string(*value)
	case *NumberType:
		return string(*value)
	case nil:
		return ""
	default:
		log.Fatal("Type not supported")
	}

	return ""
}

func GetIntPointer(value int) *int {
	return &value
}

func GetStringPointer(value string) *string {
	return &value
}
