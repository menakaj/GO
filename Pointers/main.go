package main

import (
	"fmt"
)

func add(total *int, a int, b int) {
	*total = a + b
}

func main() {
	var count1 int
	count2 := new(int)

	a := 10
	b := 30

	var x = uint8(255)
	var y = uint8(127)

	var z = x &^ y

	fmt.Println(z)

	countTemp := 10
	count2 = &countTemp
	count3 := &countTemp

	const q = 8
	const w complex128 = q
	fmt.Println(q, w)

	add(&count1, 2, 4)

	swap(&a, &b)

	fmt.Printf("A: %#v B: %#v\n", a, b)

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("Count2: %#v\n", count2)
	fmt.Printf("Count3: %#v\n", *count3)

	var int string = "as"
	var b int = 1

}
