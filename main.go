package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	var returnRate float64
	var years float64

	fmt.Print("How much you want to invest?")
	fmt.Scan(&investmentAmount)

	fmt.Print("How much your return expected?")
	fmt.Scan(&returnRate)

	fmt.Print("How long you want to invest?")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+returnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("future value: %v\nfuture value (after inflation): %v", futureValue, futureRealValue)
}
