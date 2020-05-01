package funcs

import (
	"errors"
	"fmt"
	"log"
)

func init() {
	fmt.Println("Initializing Functions.....")
}

// Add function
func Add(x, y int) int {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Fatal(errors.New("dfdsds"))
	return x + y
}

// Mul : Multiply functions
func Mul(x, y int) int {
	return x * y
}

// Div divide function
func Div(x, y int) int {

	return (x / y)
}
