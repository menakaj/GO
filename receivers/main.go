package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) Move(from int, to int) (int, int) {
	p.x = from + p.x
	p.y = from + p.y
	return p.x, p.y
}

func (p *Point) MoveP(from int, to int) (int, int) {
	p.x = from + p.x
	p.y = from + p.y
	return p.x, p.y
}

func main() {

	p := new(Point)
	p.x = 10
	p.y = 20

	p.Move(34, 23)
	fmt.Println(*p)
	p.MoveP(99, 13)
	fmt.Println(*p)

}
