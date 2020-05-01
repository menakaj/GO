package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "string", time.Now()
}

func main() {

	debug, name, t := getConfig()

	var (
		abc   bool
		fname string
		t2    time.Time
	)

	abc, fname, t2 = getConfig()

	fmt.Println(debug, name, t)
	fmt.Println(abc, fname, t2)
}
