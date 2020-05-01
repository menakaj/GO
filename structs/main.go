package main

import "fmt"

type user struct {
	id      int
	name    string
	address string
}

func checkStruct() (bool, bool) {

	u1 := user{
		id:      1,
		name:    "Menaka",
		address: "add1",
	}

	u2 := struct {
		id      int
		name    string
		address string
	}{}

	u2.id = 1
	u2.name = "Menaka"
	u2.address = "add1"

	u3 := struct {
		id      int
		name    string
		address string
	}{
		2,
		"Menaka",
		"add1",
	}

	return u3 == u2, u1 == u2
}

func main() {

	b1, b2 := checkStruct()

	fmt.Println(b1, b2)

	user2 := struct {
		add1 string
		add2 string
	}{
		"aadasd",
		"asdasdsad",
	}

	user1 := user{
		id:      1,
		name:    "tom",
		address: "232, Nw, sa",
	}

	user1.id = 2

	fmt.Println(user1, user2)
}
