package main

import "fmt"

func main() {
	per := person{"Menaka", 23, "Married"}

	fmt.Println(per.Speak())
	fmt.Println(per)
	makeSpeak(per)
}

func makeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}
