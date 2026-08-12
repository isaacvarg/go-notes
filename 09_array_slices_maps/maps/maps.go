// Package examplemaps
package examplemaps

import "fmt"

type typeMap map[string]float64

func (m typeMap) output() {
	fmt.Println(m)
}

func MapsExample() {
	// maps are always dynamic
	// so this is saying the keys and values are strings
	// you can initialize it but alsos set it to empty
	websites := map[string]string{
		"bgg": "https://boardgamegeek.com",
		"bga": "https://boardgamearena.com",
	}
	fmt.Println(websites)
	// accessing is easy
	fmt.Println(websites["bgg"])

	// since maps are dynamic you can add get accessing something that dne
	websites["chess"] = "https://chess.com"

	// delete
	delete(websites, "bga")

	// example of make with maps
	bookRating := make(map[string]int, 4)
	// this preallocates memory for a map of length 4
	bookRating["Way of Kings"] = 5

	// example of a type alias in action
	// defined typeMap above
	// this also makes the code easier to read
	anotherMap := make(typeMap, 4)
	anotherMap["hey"] = 4.44
	// now we can use that method  attached to the alias
	anotherMap.output()

	// looping arrays slices and maps
	// if don't care about index/value can also simply do
	// for range websites
	for index, value := range websites {
		fmt.Println("Index: ", index)
		fmt.Println("value", value)
	}
}
