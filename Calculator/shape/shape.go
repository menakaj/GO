package shape

import (
	"fmt"

	"../funcs"
)

func init() {
	fmt.Println("Initializing Shapes.....")
}

// Shape interface
type Shape interface {
	Area() int
}

type Square struct {
	Length int
}

type Rectangle struct {
	Length int
	Width  int
}

func (s Square) Area() int {
	return funcs.Mul(s.Length, s.Length)
}

func (s Rectangle) Area() int {
	return funcs.Mul(s.Length, s.Width)
}
