package main

import "fmt"

func main() {

	switch i := 10 % 3; i * 3 {
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Even")
	}
}
