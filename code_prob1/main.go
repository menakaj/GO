package main

import "fmt"

var maze = [4][4]int{{0, 0, 0, 0}, {1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 0}}
var order = 4
var ei = 0
var ej = 0
var steps = 0

func main() {
	start := []int{3, 0}
	fmt.Println(findPath(start[0], start[1]))

}

func findPath(si int, sj int) int {

	if si == ei && sj == ej {
		return 1
	}
	if si < 0 || sj < 0 || si == order || sj == order {
		return 0
	}
	if si+1 < order && maze[si+1][sj] != 1 && maze[si+1][sj] != 2 {
		fmt.Println("si+1 is true", si+1)
		maze[si][sj] = 2
		return 1 + findPath(si+1, sj)
	}
	if si-1 >= 0 && maze[si-1][sj] != 1 && maze[si-1][sj] != 2 {
		fmt.Println("si-1 is true", si-1)
		maze[si][sj] = 2
		return 1 + findPath(si-1, sj)
	}

	if sj+1 < order && maze[si][sj+1] != 1 && maze[si][sj+1] != 2 {
		fmt.Println("sj+1 is true", sj+1)
		maze[si][sj] = 2
		return 1 + findPath(si, sj+1)
	}

	if sj-1 >= 0 && maze[si][sj-1] != 1 && maze[si][sj-1] != 2 {
		fmt.Println("sj-1 is true", sj-1)
		maze[si][sj] = 2
		return 1 + findPath(si, sj-1)
	}
	return 0
}
