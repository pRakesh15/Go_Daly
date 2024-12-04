package main

import "fmt"

func printStar(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("this is a main function")
	var row int
	fmt.Scan(&row)
	printStar(row)
}
