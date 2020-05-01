package main

type student struct {
	StudentId int `json:"id"`
	LastName string `json:"lName"`
	MiddleInitial string `json:"minitial"`
	FirstName string `json:"fName"`
	IsEnrolled bool `json:"enrolled"`
	Courses []course `json:"classes"`
}

type course struct {
	Name string `json:"cName"`
	Nubmer int `json:"cNo"`
	Hours int `json:"cHours"`
}