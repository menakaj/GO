package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name        string
	age         int
	civilStatus string
}

func (p person) Speak() string {
	return p.name
}

func (p person) String() string {
	return fmt.Sprintf("Person Name %s, Age %d, Civil Status %s", p.name, p.age, p.civilStatus)
}
