package main

import "fmt"

func a() error {
	return nil
}

func main() {

	anon := func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
		fmt.Println("this is deferred function")
	}

	defer anon()

	_ = a()

	// fmt.Println(GetError(210))

	fmt.Errorf("After panic %s", "sss")
}
