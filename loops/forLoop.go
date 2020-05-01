package main

import "fmt"

func forLoop() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func infinite() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i > 1000 {
			break
		}
	}
}

func loopArray() {
	arr := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func loopMap() {
	config := map[int]string{
		1: "asda",
		2: "sdad",
		3: "sfd",
	}

	for key, val := range config {
		fmt.Println(key, val)
	}
}

func findMax() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	max := 0
	var word string

	for key, value := range words {
		if value > max {
			max = value
			word = key
		}
	}
	fmt.Println(word, max)
}
