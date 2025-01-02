package main

import (
	"fmt"
)

func main() {

	//start defer key word which is a package of main
	fmt.Println("start of the programmer")
	defer fmt.Println("middle  of the programmer")     //after execute  of the function it will print
	defer fmt.Println("2nd middle  of the programmer") //after execute  of the function it will print it is the last defer so it will execute  1st..
	fmt.Println("end  of the programmer")
	// "2nd middle  of the programmer"
	// "middle  of the programmer"

	//outPut is
	// 	start of the programmer
	// end  of the programmer
	// middle  of the programmer

	//after the execution of the function the defer is execute...
	//in the adobe example when the main function is complete execution then  the defer start execute .. so  "middle  of the programmer" comes  to last

	//if there are 2 defer in a single function
	//it work like a queue and the principle of LIFO("last in fast out")
	// outPut is
	// 	start of the programmer
	// end  of the programmer
	// 2nd middle  of the programmer
	// middle  of the programmer
}
