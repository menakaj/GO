package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
	Debug2      bool
	Logs        = "debug"
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime, Debug2, Logs)
}
