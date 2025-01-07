package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Print("error is ", err)

		return
	}

	fmt.Printf("type of response is %T", res)
	// fmt.Println("the  response is", res)

	data, errors := ioutil.ReadAll(res.Body)

	if errors != nil {
		fmt.Print("error is ", errors)

		return
	}
	obj := string(data)
	fmt.Println("data is ", obj)

}
