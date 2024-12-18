package main

import "fmt"

func main() {
	//here i declare a map name as m ..
	//as we know map has key value pair here the type is string and ing
	//string for key and int for value...
	// inictilazi the map 1st and then declare the key value..
	m := make(map[string]int)

	//declare the  key value pair....
	m["Rakesh"] = 23
	m["Bikash"] = 29
	m["Prakesh"] = 28

	//print the type of m which is map[].

	fmt.Printf("type of  the variable is %T \n", m)
	fmt.Println(m["Bikash"])

	//checking  of presence of a key inside a map
	//create and declare the map at a time..
	mark := map[string]int{"Rakesh": 25, "Subhasis": 27, "Mithun": 29, "Lipun": 36}

	//by using %T we  can determine the type of any variable
	fmt.Printf("type of mark is %T \n", mark)
	//len function for find the length od array,slice,map and many more...
	fmt.Println("length  of mark is ", len(mark))

	//declare some key value pairs..
	// mark["Rakesh"] = 25
	// mark["Subhasis"] = 29
	// mark["Mithun"] = 30

	//checking  individual key..
	//by default if there is no key match it return 0
	fmt.Println("mark of mithun is ", mark["Mithun"])

	//deleting specific key and corresponding value
	delete(mark, "Mithun")
	fmt.Println("mark of mithun is ", mark["Mithun"])

	//im map when we try to fetch a key it will return 2 things
	// one is the value of that key and another one is a bool(true or false) value of the key is present or not

	value, exist := mark["Rakesh"]
	fmt.Println("rakesh is present or not", exist) //return tru
	fmt.Println("mark of rakesh is ", value)       //return value=25

	value1, exist1 := mark["Mithun"]

	if exist1 {
		fmt.Println("Mithun bhai's mark is ", value1)
	} else {
		fmt.Println("Mithun bhai is not present")
		fmt.Println(exist1)

	}

	//applying for loop for read all the key value pair in the mark map

	for index, value := range m {
		fmt.Printf("name is %s and the age is %d  \n", index, value)
	}

}
