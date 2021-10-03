package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe(":8090", nil)
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello received....")
	fmt.Println(req.Header)
}
