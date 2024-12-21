package main

import (
	"fmt"
	// "strconv"
)

//As we know pointer is also a variable but it not store value
// it stores the address of another variable.
//it provides a way to direct interact with the memory..

// pass a variable as reference or pointer..
func pointerFunction(num *int) {

	//i am not returning any  thing bcz it will directly modify the value in memory...
	*num += 10
}

func main() {
	//here i simple declare a variable
	var number int = 32
	fmt.Println(&number) //0xc00000a0b8

	//declare a pointer

	// var ptr *int

	//in the above pointer declaration the int is note for the value of the pointer or the address of the memory
	//but it is for the variable assigned to this pointer is holding the value is type integer.
	//assign or point the pointer to number variable
	ptr := &number //& means address of the variable.

	fmt.Println(ptr) //0xc00000a0b8  this  is the address of number variable

	//want to print  the data contain in the variable through the pointer ..
	fmt.Println("data through the pointer is ", *ptr)

	var pointer *string

	fmt.Print(pointer) //print the default value of pointer that is ==> <nil>

	value := 30
	pointerFunction(&value)

	fmt.Print(value)

	///lets work pointer with array

	arr := []int{1, 5, 6, 8}

	fmt.Print("\n", &arr[0]) //0xc000016200
	fmt.Print("\n", &arr[1]) //0xc000016208
	fmt.Print("\n", &arr[2]) //0xc000016210
	fmt.Print("\n", &arr[3]) //0xc000016218
	// fmt.Print("\n", &arr[4]) ////runtime error: index out of range [4] with length 4

	// num := 123

	// str := strconv.Itoa(num)
	// str += "rakesh"

	// fmt.Print(str)

}
