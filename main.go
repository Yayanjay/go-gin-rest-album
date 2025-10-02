package main

import (
	"fmt"
	"golang-rest-api/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error getting account balance")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to Go Bank")
	for i := 0; i < 10; i++ {
		menuOptions()

		var input int
		fmt.Print("your choice is ")
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated, your current balance is: ", accountBalance)
		default:
			fmt.Println("See ya")
			fmt.Println("Thanks for using our bank")
			return
		}
	}

}
