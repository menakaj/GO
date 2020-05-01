package main

import "fmt"

func main() {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	days2 := []string{"Sun"}

	days2 = append(days2, days[0:6]...)

	fmt.Println(days2)
}
