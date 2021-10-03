package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := []byte(`
	{
		"fname" : "Menakaa"
	}
	`)

	studentInfo := []byte(
		`{ 
			"id": 123, 
			"lname": "Smith", 
			"minitial": null, 
			"fname": "John", 
			"enrolled": true, 
			"classes": [{ 
			  "coursename": "Intro to Golang", 
			  "coursenum": 101, 
			  "coursehours": 4 
			}, 
			{ 
			  "coursename": "English Lit", 
			  "coursenum": 101, 
			  "coursehours": 3 
			}, 
			{ 
			  "coursename": "World History", 
			  "coursenum": 101, 
			  "coursehours": 3 
			} 
		  ] 
		   }`)

	type course struct {
		Name   string `json:"coursename"`
		Nubmer int    `json:"coursenum"`
		Hours  int    `json:"coursehours"`
		Marks int 	   `json:"marks,omitempty",default: "2"`
	}

	type student struct {
		StudentId     int      `json:"id"`
		LastName      string   `json:"lname"`
		MiddleInitial string   `json:"minitial"`
		FirstName     string   `json:"fname"`
		IsEnrolled    bool     `json:"enrolled"`
		Courses       []course `json:"classes"`
	}

	type person struct {
		Name string `json:"fname"`
		ID string `json: "id, omitempty"`
	}

	var stu student

	 p := person{Name: "Abcs", ID:"dfsdfd"}

	err := json.Unmarshal(data, &p)

	p2 := person{Name: "Saman", ID:"Saman"}

	err2 := json.Unmarshal(studentInfo, &stu)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(stu)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}

	byts, err := json.Marshal(p2)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byts))
}
