package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "current.txt"

func readFromFile() float64 {
	// the _ tells go you know there is a value here but
	// don't want to deal with it rn
	data, _ := os.ReadFile(fileName)
	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)
	return value
}

func writeToFile(currentNumber float64) {
	// helps convert this to a formatted string which
	// can then be converted to bytes using []byte()
	convertedBytes := fmt.Sprint(currentNumber)
	os.WriteFile(fileName, []byte(convertedBytes), 0644)
}

func main() {
	currentNumber := readFromFile()

	fmt.Println(currentNumber)

	for {

		fmt.Println("1 to add 10")

		fmt.Println("2 to read number")

		var choice int

		fmt.Print("What we do?")
		fmt.Scan(&choice)

		switch choice {
		case 1:

			fmt.Println("hello")
		case 2:
			fmt.Println("hey")
		default:
			return
		}
	}
}
