package main

import "errors"

func addWithSwitch(a interface{}, b interface{}) (interface{}, error) {
	switch val := a.(type) {
	case int:
		return val, nil
	case bool:
		return "bool", nil
	case string:
	default:
		return "string", nil
	}

	return "", errors.New("Unsupported type")
}
