package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) float64 {
	return float64(x) / float64(y)
}

func main() {
	fmt.Println("************ Welcome to Calculator")

	var x, y int

	fmt.Println("enter 1st no")
	fmt.Scan(&x)
	fmt.Println("enter 2nd no")
	fmt.Scan(&y)

	fmt.Println("for addition press 1")
	fmt.Println("for subtraction  press 2")
	fmt.Println("for multiplication  press 3")
	fmt.Println("for deviation  press 4")
	fmt.Println("for exit  press 5")

	var option int

	fmt.Scan(&option)

	switch option {
	case 1:
		var ans int = add(x, y)
		fmt.Println("your ans is ", ans)
		main()
	case 2:
		var ans int = sub(x, y)
		fmt.Println("your ans is ", ans)
        main()
	case 3:
		var ans int = mul(x, y)
		fmt.Println("your ans is ", ans)
		main()
	case 4:
		var ans float64 = div(x, y)
		fmt.Println("your ans is ", ans)
        main()
	case 5:
		fmt.Println("thank you for visiting this calculator...")
	default:
		fmt.Println("Opps your are note selecting the given option")
		main()

	}

}
