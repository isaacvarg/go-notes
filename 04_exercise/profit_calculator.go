package main

import (
	"errors"
	"fmt"
	"os"
)

// goal
// validate user input
// 		- show error message and exit if invalid input provided
// 		- no negavitive numbers
// 		- no zero
// store calculated results in file

func main() {
	revenue, errorRev := getUserInput("Revenue: ")
	expenses, errExp := getUserInput("Expenses: ")
	taxRate, errRat := getUserInput("Tax Rate: ")

	if errorRev != nil || errExp != nil || errRat != nil {
		// they error message is the same
		panic(errorRev)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	saveCalc(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func saveCalc(ebt, profit, ratio float64) {
	converted := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f", ebt, profit, ratio)
	err := os.WriteFile("calc.txt", []byte(converted), 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)

	if userInput == 0 || userInput < 0 {
		return 0, errors.New("inputs must be non-zero and positive values")
	}

	if err != nil {
		fmt.Println(err)
	}
	return userInput, nil
}
