package main

import "fmt"

type One_o struct {
	id int
}

type One struct {
	Name string
	Id   One_o
}

type Abc struct {
	Items []One
}

func main() {

	one0 := One_o{
		id: 10,
	}

	one2 := One_o{
		id: 20,
	}

	oneOne := One{
		Name: "this is one one",
		Id:   one0,
	}
	oneTwo := One{
		Name: "This is one two",
		Id:   one2,
	}

	ab3 := []Abc{
		Abc{
			Items: []One{oneOne, oneTwo},
		},
		Abc{
			Items: []One{oneOne},
		}}

	fmt.Println(ab3)

}
