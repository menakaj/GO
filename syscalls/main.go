package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for {
			s := <-sigs

			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("Process is being interrupted...")
				done <- true
			}
		}
	}()
	fmt.Println("Blocked...")
	<-done
}
