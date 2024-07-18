package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.0
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do? ")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ", choice)
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("How much would you like to deposit?")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			} else if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Println("Your new balanace is: ", accountBalance)
			}
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("How much would you like to withdraw? ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than you have.")
				continue
			} else if withdrawAmount > 0 {
				accountBalance -= withdrawAmount
				fmt.Println("Your new balanace is: ", accountBalance)
			}
		} else {
			println("Thank you for using Go Bank. Goodbye!")
			break
		}
	}

}
