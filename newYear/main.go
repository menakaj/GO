package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseUint("-111", 10, 32))
}
