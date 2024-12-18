package main

import "fmt"

// the structure (struct) in Go (Golang) is conceptually similar to an object in JavaScript

//declare a structure out side the function

type Address struct {
	state string
	dist  string
	lain  string
	pin   int32
}

type bioData struct {
	name    string
	age     int
	mobile  int64
	address string
}

//here we can assign  a structure inside a structure...

type BioData struct {
	name    string
	age     int
	mobile  int64
	address Address
}

func main() {

	//here we use structure of bioData which is declare outside the main function
	var person1 bioData

	//there is a another method to declare a variable  by using new key word
	//here new key create a memory and pointing  to person
	//here person is a kind of pointer
	var person = new(BioData)

	person.name = "Piush"
	person.age = 23
	person.mobile = 9876528563
	// person.address.dist="Puri"
	// person.address.state="Odisha"
	// person.address.lain="Reanghalo"
	// person.address.pin=752114
	person.address = Address{state: "Odisha", dist: "Puri", lain: "Renghalo", pin: 752114}
	

	fmt.Println("person is  ", person) //&{Piush 23 9876528563 BadaMachpur}
	fmt.Println("person is from ", person.address.state)

	//add the value for person1 in a way where 1st declare then add data
	person1.name = "Rakesh Pradhan"
	person1.age = 23
	person1.mobile = 6372700872
	person1.address = "konark"

	//there is a another method to do the same
	//we can derict  create and add data

	person2 := bioData{name: "Sarthak Rout", age: 23, mobile: 7539514562, address: "Renghlo"}

	//print the total person
	fmt.Print(person2)

	//print specific  function
	fmt.Println(person1.name)
}
