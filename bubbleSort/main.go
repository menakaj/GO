package main

import "fmt"

func swap(a int, b int) {

}

func main() {
	arr := []int{3, 5, 2, 6, 2, 1, 6, 8, 9, 1, 0, 4, 7, 34, 46, 23, 525, 5, 345, 35234, 23423, 432, 423, 423, 423, 4, 23, 423, 4, 234, 324, 23, 423, 4}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}
		if arr[i] < arr[i-1] {
			for j := i; j > 0; j-- {
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}

			}
		}
	}
	fmt.Println(arr)
}
