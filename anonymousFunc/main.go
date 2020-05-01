package main

import "fmt"

type calc func(int, int) string

func main() {

	defer func() {
		fmt.Println("After all")
	}()
	fmt.Println(salary(100, 200, empSal))
	fmt.Println(salary(20, 332, manSal))
}

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func empSal(a, b int) int {
	return a * b
}

func manSal(a, b int) int {
	return a + b
}
