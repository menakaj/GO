package main

import (
	"fmt"
	"sync"
)

func gen(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range num {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(ch chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(ch))

	for _, c := range ch {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	//nums := []int{3, 2, 3, 4, 1, 1, 5, 3}
	c1 := gen(1, 2, 3, 4, 5)
	sq1 := sq(c1)
	sq2 := sq(c1)
	// sq3 := sq(c1)
	// sq4 := sq(c1)

	for n := range merge(sq1, sq2) {
		fmt.Println(n)
	}

}
