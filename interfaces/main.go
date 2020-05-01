package main

import (
	"errors"
	"fmt"
)

func add(a interface{}, b interface{}) (interface{}, error) {

	if a, e := a.(int); e {
		if b, e2 := b.(int); e2 {
			return a + b, nil
		}
	}

	if a, e := a.(string); e {
		if b, e2 := b.(string); e2 {
			return a + " " + b, nil
		}
	}
	return "", errors.New("Un SUpported type")
}

func main() {
	fmt.Println(add(3, 4))

	fmt.Println(add(2, "sadas"))

	fmt.Println(addWithSwitch("dfsdf", 3))

	fmt.Println(checkType(1 == 1))

	ca := cat{}
	fmt.Println(ca.Speak("sasdds"))
	fmt.Println(ca.Move(2, 4))
	fmt.Println(ca)

}
