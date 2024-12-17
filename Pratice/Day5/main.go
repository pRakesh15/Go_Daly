package main

import "fmt"

//this is  wrong because this type declare is not allowed outside the function or method..
// age := 23

//this is a constant type and this is not exportable like this is not public 
//or not accessible out side this main file 
const password int=15852


//this is constant and is  also public bcz the variable starts from capital letter
const Pin int=1459

func main() {
	//this is a type we can declare a variable
	var myName string = "Rakesh Pradhan"

	//this is also  a type by which we can declare a variable but there is a problem
	//we only can declare this type of variable inside a method or a  function
	name := "Rakesh Pradhan"

	fmt.Println(name)
	fmt.Println(myName)
	fmt.Println(password)
}
