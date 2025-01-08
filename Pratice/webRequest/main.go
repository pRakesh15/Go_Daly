package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Comp   bool   `json:"comp"`
}

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Print("error is ", err)

		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error while gating the response", res.Status)

		return
	}

	// fmt.Println("status code is ", res.StatusCode)
	// fmt.Println("the  response is", res.Body)

	// data, errors := ioutil.ReadAll(res.Body)

	// if errors != nil {
	// 	fmt.Print("error is ", errors)

	// 	return
	// }
	// fmt.Println("data is ", string(data))

	var todo Todo

	errr := json.NewDecoder(res.Body).Decode(&todo)

	if errr != nil {
		fmt.Println("the error  is ", errr)

		return
	}

	fmt.Println("the data is ", todo.Comp)

}
