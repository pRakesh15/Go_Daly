package main

import (
	"fmt"
)

// func printMyname() {
// 	fmt.Println("hy my name is rakesh")
// }

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Jay shree RAM")
	// var name string = "Rakesh"
	 n := 15
	fmt.Println(n)
	var x = Person{
		name: "rakesh",
		age:  20,
	}
	x.name = "Bikash"
	x.age = 29

	myname := "Rakesh"
	myage := 23

	fmt.Println(x)

	fmt.Printf("my name is %s and age is  %d", myname, myage)
	// fmt.Printf("my age is %d", myage)

}
