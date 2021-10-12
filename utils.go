package main

import (
	"log"
	"strconv"
)

// ValueToString convert a value to string
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

// GetIntPointer get a int pointer to value
func GetIntPointer(value int) *int {
	return &value
}
