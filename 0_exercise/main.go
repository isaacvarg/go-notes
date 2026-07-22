package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("Enter values separated by space:")
	fmt.Print("<Revenue> <Expenses> <taxRate>: ")
	item, err := fmt.Scan(&revenue, &expenses, &taxRate)

	fmt.Println(item, err)

	ebt := revenue - expenses
	afterTax := ebt * (taxRate / 100)
	profit := ebt - afterTax
	ratio := ebt / profit

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("EBT/Profit Ratio: ", ratio)
}
