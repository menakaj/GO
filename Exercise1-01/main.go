package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)

	var fName string = "Menaka"
	var address string = "dfsdf"
	var age int = 23

	fmt.Println(fName, address, age)

}
