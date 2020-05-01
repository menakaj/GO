package main

import "fmt"

func swap(a *int, b *int) {
	// swap the values here
	fmt.Printf("Before A : %#v B : %#v\n", *a, b)
	A := *a

	*a = *b
	*b = A

	fmt.Printf("After A : %#v B : %#v\n", *a, *b)

}
