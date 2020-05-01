package main

import (
	"fmt"
)

func swap(arr []int, x int, y int) []int {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
	return arr
}

func main() {
	arr := []int{0, 3, 3, -6, 3, -3, 1}
	removeElement()

	i, j, k := 0, 0, len(arr)-1
	for j <= k {
		if arr[j] < 0 {
			arr = swap(arr, i, j)
			fmt.Println("<0 i:", i, "j:", j, "k:", k, arr)
			j++
			i++
		} else if arr[j] == 0 {
			j++
			fmt.Println("=0 i:", i, "j:", j, "k:", k, arr)
		} else {
			arr = swap(arr, j, k)
			k--
			fmt.Println(">0 i:", i, "j:", j, "k:", k, arr)
		}
	}
	fmt.Println(arr)
}
