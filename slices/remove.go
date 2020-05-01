package main

import "fmt"

func removeElement() {
	items := []string{"Good", "Good", "Bad", "Good", "Good"}

	itemsDeleted := append(items[:2], items[3:]...)

	fmt.Println(itemsDeleted)
}
