package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "myFile.txt"

// this old function doesn't crash if the file dne
// in go functions are written such that they
// don't crash the code, but return null values so that
// the application doesn't crash
// such as below, the file dne, it parses an empty data var
// then this is converted to 0 from the ParseFloat
//
// func readFromFile() float64 {
// 	// the _ tells go you know there is a value here but
// 	// don't want to deal with it rn
// 	data, _ := os.ReadFile(fileName)
// 	valueText := string(data)
// 	value, _ := strconv.ParseFloat(valueText, 64)
// 	return value
// }

// error is type built into go
func readFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	// nil in go is special valure that means absense of useful value
	if err != nil {
		return 0, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to convert value")
	}
	return value, nil
}

func main() {
	fileContent, err := readFromFile()
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		// could do a return here to exit in
		// cases where code cannot continue with that error
		// but panic is a better way to exit and display message
		panic("Can't fun no more!")
	}

	fmt.Println(fileContent)
}
