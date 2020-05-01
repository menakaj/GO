package main

import (
	"fmt"
)

func main() {
	count := 10

	fmt.Println(count)

	count++
	fmt.Println("Count++ ", count)

	count += 100
	fmt.Println("Count + 100 ", count)

	name := "Menaka"

	name += " Jayawardena"

	fmt.Println(name)
}
