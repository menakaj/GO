package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func checkPassword(pwd string) bool {
	pWD := []rune(pwd)

	hasNumbers := false
	hasUpper := false
	hasSybol := false

	b := []byte(pwd)
	fmt.Println("Password: ", b)
	fmt.Println("Rune: ", pWD)

	if len(pWD) < 8 {
		return false
	}

	for _, char := range pWD {
		if unicode.IsNumber(char) {
			hasNumbers = true
			continue
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSybol = true
			continue
		}
		if unicode.IsUpper(char) {
			hasUpper = true
			continue
		}
	}
	return hasNumbers && hasUpper && hasSybol
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Password --> ")

	text, _ := reader.ReadString('\n')
	fmt.Println(checkPassword(text))
}
