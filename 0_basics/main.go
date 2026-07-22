package main

// used `go mod init github.com/isaacvarg/go-bacis` to create module
// `go run .` for dev run

import (
	"fmt"
	"math"
)

func main() {
	var investment float64 = 1000

	// this declaration is for vars with infered type
	rate := 5.5
	var years float64 = 10

	value := investment * math.Pow(1+rate/100, years)

	// can also declare multiple vars at once
	// if set a type, both values must be that type
	// infered values can be multiple types
	first, last := "isaac", "vargas"

	// while that feature is cool, i feel like a lot of variables makes the code harder to read.
	// I will have to play around with it

	// const similar to typescript, can not change
	const coolNumber float64 = 200.5

	fmt.Println(first, last)
	fmt.Println(value)
	fmt.Println(coolNumber)

	// Scan method allows user input
	// single vaue limitation for scan
	// `&` is a pointer to that var
	var greeting string
	fmt.Print("Enter a greeting: ")
	// i was getting a syntax error to capture these
	// not covered in my learning so far but cool that it ensures you capture
	n, err := fmt.Scan(&greeting)
	fmt.Print(n, err)
	fmt.Println(greeting)
}
