package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args
	names := map[string]string{
		"112": "menaka",
		"321": "Chrysh",
		"01":  "nimal",
	}

	if len(input) < 2 {
		os.Exit(0)
	}

	for k, v := range names {
		if k == input[1] {
			fmt.Println("Hi ", v)
		}
	}

	abc := append(input, os.Args...)
	fmt.Println(abc)
}
