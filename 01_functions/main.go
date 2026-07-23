package main

import "fmt"

func main() {
	total, randomNumber := calculateThis(20.6, 22.0)
	fmt.Println(total, randomNumber)

	sum := addOne(20.0)
	fmt.Println(sum)
}

// two parameters
// outputs specified
func calculateThis(numA, numB float64) (float64, float64) {
	total := numA * numB
	anotherVariable := 16.0

	// can return multiple things directly
	return total, anotherVariable
}

// alternate syntax
// this initliazed a null variable (sum)
// then when using the return it automatically returns that var
// this is cool, but I am so used to more explicit declaration of the return
// it is just easier for me to read, but i might get used to this
func addOne(numA float64) (sum float64) {
	sum = numA + 1
	return
}
