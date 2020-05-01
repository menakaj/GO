package main

import "fmt"

func nonLoggedHours() func(int) int {
	hoursNotLogged := 0
	return func(x int) int {
		hoursNotLogged += x
		return hoursNotLogged
	}
}

// func PayDay() (int, bool) {

// }

func payDetails() {

}

func main() {
	notLogged := nonLoggedHours()

	fmt.Println(notLogged(2))
	fmt.Println(notLogged(3))
	fmt.Println(notLogged(5))

}
