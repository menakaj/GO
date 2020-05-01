package main

import "fmt"

func checkType(param interface{}) string {
	switch val := param.(type) {
	case int:
		fmt.Println(val)
		return "is Integer"
	case bool:
		return "Boolean"
	case float32:
	case float64:
		return "float"
	case string:
		return "string"
	default:
		return "unknown"
	}
	return "Unkown"
}
