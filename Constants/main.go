package main

import "fmt"

const GlobalLimit = 1000
const MaxCacheSize = 10 * GlobalLimit

const (
	CachaKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func addToCache(key string, name string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = name
}

func setBook(isbn string, name string) {
	addToCache(CachaKeyBook+isbn, name)
}

func setCD(cd string, name string) {

	addToCache(CacheKeyCD+cd, name)
}

func getBook(isbn string) string {
	return cache[CachaKeyBook+isbn]
}

func getCD(cd string) string {
	return cache[CacheKeyCD+cd]
}

const (
	_      = iota
	Sunday = 1 << (10 * iota)
	TuesDay
)

func main() {
	cache = make(map[string]string)
	setBook("1234", "fsdfsdfds")
	setCD("asd", "asdqwerew")

	fmt.Println(getBook("1234"))
	fmt.Println(getCD("asd"))
	fmt.Println(1 << 10 * 2)
}
