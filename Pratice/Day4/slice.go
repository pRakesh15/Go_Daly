package main

import "fmt"

func main() {

	//this is a slice because the length of the array is dynamic
	// in this example we can use random no of elements inside the number slice...
	number := []int{1, 5, 8, 2, 3, 5, 8, 9, 7, 8}
	number = append(number, 1, 58, 96, 85, 7, 58, 785)
	var x string

	fmt.Println(number)
	fmt.Printf("%T", x)
}
