package main

import (
	"encoding/gob"
	"os"
	"unsafe"
)

func main() {
	da2 := []string{"12213wwww", "12312", "12312"}

	files, _ := os.Create("db.text")

	enc := gob.NewEncoder(files)
	enc.Encode(da2)
	files.Close()

	var data []string
	f2, _ := os.Open("db.text")
	n, _ := os.Stat("db.text")
	fmt.Println(unsafe.Sizeof(n))
	fmt.Println(unsafe.Sizeof(da2))

	en2 := gob.NewDecoder(f2)

	en2.Decode(&data)
	fmt.Println(data)
	f2.Close()
}
