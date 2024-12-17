package main

//here i import a package name as fmt which is a prebuilt go package...

import "fmt"

//here the 1st int is for type of input
//and the 2nd tye bool is  for the output type like what is the type  of output this function is return....
func checkPass(currentPin, enteredPin int) bool {

	return currentPin == enteredPin
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	const Pin int = 143500
	var bankBalance int = 236
	fmt.Println("************ Welcome to SBI ATM service")

	var password int

	//there are lot of function are define inside the fmt package ....
	//like Print Scan...
	//Print for output ...
	//Scan for input ....
	//mainly FMT package is known for IO input and output  ....
	//As there are some function are import from fmt package ..
	//The function are written in Capital letter...
	//Printf...Scan.....

	fmt.Println("enter your password plz")
	fmt.Scan(&password)

	var check bool = checkPass(Pin, password)

	if check {
		fmt.Println("Welcome Mr. Rakesh Pradhan to SBI")

		fmt.Println("****** Options *********")

		fmt.Println("for check bank balance press 1 ")
		fmt.Println("for deposit press 2")
		fmt.Println("for withdraw press 3")
		fmt.Println("for exit  press 0")
		var option int

		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("your bank balance is  ", bankBalance)
			main()
		case 2:
			fmt.Println("please enter deposite amount")
			var amount int

			fmt.Scan(&amount)
			var totalAmount int = add(bankBalance, amount)
			fmt.Println("Deposit successful your Current bank balance is  ", totalAmount)
			main()
		case 3:
			fmt.Println("please enter withdraw  amount")
			var amount int
			fmt.Scan(&amount)
			if amount > bankBalance {
				fmt.Println("insufficient found ")
				main()
			} else {
				var totalAmount int = sub(bankBalance, amount)
				fmt.Println("withdrawal  successful your Current bank balance is  ", totalAmount)
				main()
			}

		case 0:
			fmt.Println("thank you for visiting this calculator...")
		default:
			fmt.Println("Opps your are note selecting the given option")
			main()

		}
	} else {
		fmt.Println("incorrect password")
	}

}
