package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, req *http.Request, params router.Params) {
	fmt.Println("Index", params)
}

func PageOne(w http.ResponseWriter, req *http.Request, params router.Params) {
	fmt.Println("Page One", params)
}

func getContent(w http.ResponseWriter, req *http.Request, params router.Params) {
	fmt.Println("GEt content....", params)
}

func main() {
	r := router.New()

	r.GET("/api/:id/content", getContent)
	r.GET("/", Index)
	r.POST("/page/:number", PageOne)

	log.Fatal(http.ListenAndServe(":8080", r))
}
