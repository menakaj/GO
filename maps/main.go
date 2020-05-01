package main

import "fmt"

func main() {
	names := map[string]string{}

	names["abc"] = "ancss"
	names["abc2"] = "ancss"
	names["abc4"] = "ancss"

	fmt.Println(names)

	delete(names, "abc")

	fmt.Println(names)

}
