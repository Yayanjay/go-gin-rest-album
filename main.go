package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("How much your revenue?")
	fmt.Scan(&revenue)

	fmt.Print("How much your expenses?")
	fmt.Scan(&expenses)

	fmt.Print("How much your tax rate?")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax
	fmt.Println("EBT =", earningBeforeTax)
	fmt.Println("EAT =", earningAfterTax)
	fmt.Println("RATION =", ratio)
}
