package main

import (
	"fmt"

	. "./funcs"
	"./shape"
)

func init() {
	count++
	fmt.Println("Initializing Main.....")
}

func abc() int {
	return 10
}

var count = abc()

func main() {
	rect := shape.Rectangle{Length: 3, Width: 5}
	sq := shape.Square{Length: 23}

	defer func() {
		fmt.Println(count)
	}()

	totalArea := Add(rect.Area(), sq.Area())
	fmt.Println(totalArea)
}
