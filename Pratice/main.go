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

	fmt.Println("my name is %s and age is  %d", myname, myage)
	// fmt.Printf("my age is %d", myage)

	//code for read and write files

	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Print("error while creating file", err)
	// 	return
	// }

	// content := "hy my name is rakesh"

	// _, errors := io.WriteString(file, content)

	// if errors != nil {
	// 	fmt.Println("error while writing a file", errors)

	// 	return
	// }

	// fmt.Print("file created successfully")

	// defer file.Close()

}
