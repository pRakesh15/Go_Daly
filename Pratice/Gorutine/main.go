package main

import (
	"fmt"
	"time"
)

// is the 1st function
func sayHello() {

	for i := 0; i < 5; i++ {
		fmt.Println("Hello buddy")
		time.Sleep(10 * time.Millisecond)
	}

	//this is like setTimeout in js basically it will delay the function for 1sec..
	// fmt.Println("hyy after 1sec delay")
}

// is the 2nd function
func sayHii() {
	for i := 0; i < 5; i++ {
		fmt.Println("hii from Rakesh :)")

	}
	time.Sleep(10 * time.Millisecond)
}

func main() {
	fmt.Println("this  is a function for goroutine")

	//now sayHello function run in a chanel and in same chanel the sayhii is also run
	//which means 1st function1 run after complete of fun1 the fun2 start running
	//this code is in single thrade

	// sayHello()
	// sayHii()

	//lets make it concorent through goroutine
	//here we make the functions independent by  using goroutines
	//both the functions start running symontaniosly

	// these function take 1sec to run
	// go sayHello()
	// //this function execute 1st
	// sayHii()

	//this is for wt 2 for the complete of the delayed function

	go sayHello()
	sayHii()
	time.Sleep(2 * time.Second)

}
