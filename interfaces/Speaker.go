package main

type Speaker interface {
	Speak(string) string
	Move(int, int) (int, int)
}

type cat struct {
}

func (c cat) Speak(word string) string {
	return word
}

func (c cat) Move(x, y int) (int, int) {
	return x + 10, y + 10
}
