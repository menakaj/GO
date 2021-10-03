package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"

	"github.com/fsnotify/fsnotify"
)

var once sync.Once

//type sender chan string

var (
	C chan string
)

// NewSender ...
func NewSender() chan string {
	once.Do(func() {
		C = make(chan string)
	})
	return C
}

// GetReceiver ...
func GetReceiver() chan string {
	once.Do(func() {
		C = make(chan string)
	})
	return C
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		log.Println("Handler invoked ", name)
		h(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()
	c := NewSender()
	c <- "Invoked..........."
	log.Println("Start waiting....")
	sendRequest(r, w)
}

func sendRequest(r *http.Request, rw http.ResponseWriter) {
	ur, _ := url.Parse("http://localhost:3001")

	req, _ := http.NewRequest(r.Method, ur.String(), r.Body)

	req.Header.Set("X-Fowarded-For", r.RemoteAddr)
	req.Header.Set("Host", r.Host)

	for header, value := range r.Header {
		for _, val := range value {
			req.Header.Set(header, val)
		}
	}
	fmt.Println(ur)
	fmt.Println(r)

	proxy := httputil.NewSingleHostReverseProxy(ur)
	proxy.ServeHTTP(rw, r)
}

func sendAnalytics() {
	c := GetReceiver()
	fmt.Println("In receiver....")
	for d := range c {
		log.Println("Receiving....")
		log.Println(d)
	}
}

func startServer() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/a", handleGet)
	server.ListenAndServe()
}

func main() {

	watcher, _ := fsnotify.NewWatcher()
	watcher.Add("anc.text")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	//handler := MyHandler{}

	//anCh := make(chan string, 10)
	go sendAnalytics()
	go startServer()

OUTER:
	for {
		select {
		case <-sig:
			switch <-sig {
			case syscall.SIGHUP:
				fmt.Println("SIGHUP")
			case syscall.SIGABRT:
				fmt.Println("SIG ABRT")
			}
			break OUTER
		case e := <-watcher.Events:
			fmt.Println("Something happened", e)
			f, _ := os.Open("anc.text")
			b, _ := ioutil.ReadAll(f)
			fmt.Println(string(b))
		}
	}

	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 		time.Sleep(time.Second)
	// 	}
}
