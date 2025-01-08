package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("this is a main function")

	person := Person{Name: "Rakesh", Age: 24, IsAdult: true}
	// fmt.Printf("type of person is %T \n", person)
	fmt.Println("person age is ", person.Age)

	personData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error is ", err)
	}

	fmt.Println(string(personData))

}
