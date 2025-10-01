package main

import (
	"fmt"
)

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank")
	for i := 0; i < 10; i++ {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var input int
		fmt.Println("your choice is", input)
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Println("your balance is ", accountBalance)
		case 2:
			var inputDeposit float64
			fmt.Print("how much do you want to deposit? ")
			fmt.Scan(&inputDeposit)
			if inputDeposit <= 0 {
				fmt.Println("you need to deposit grater than zero")
				continue
			}
			accountBalance += inputDeposit
			fmt.Println("Balance updated, your current balance is: ", accountBalance)
		case 3:

			var inputWithdraw float64
			fmt.Print("how much do you want to withdraw? ")
			fmt.Scan(&inputWithdraw)
			if inputWithdraw <= 0 {
				fmt.Println("you need to deposit grater than zero")
				continue
			}

			if inputWithdraw > accountBalance {
				fmt.Println("you can't withdraw more than you have")
				continue
			}
			accountBalance -= inputWithdraw
			fmt.Println("Balance updated, your current balance is: ", accountBalance)
		default:
			fmt.Println("See ya")
			fmt.Println("Thanks for using our bank")
			return
		}
	}

}
