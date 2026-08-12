package main

import (
	"fmt"

	examplemaps "github.com/isaacvarg/array/maps"
)

// structus very useful because they group different types of data
// into one piece of data
// array is similar to Ts, same as structs are the interface or type
func main() {
	// can declare without assigning
	var names [4]string
	// zero indexed, like everyting else
	price := [4]float64{10.99, 7.88, 10.00, 8.88}
	fmt.Println(price)
	fmt.Println(price[2])

	// sets the third slot in the array
	names[2] = "Isaac"

	// Slices
	// gets a subset of an array
	// returns 7.88 and 10.00
	// first value included last not
	// can also use price[:3]
	// which slices up until the specified value
	// same as pridce[1:]
	featuredPrice := price[1:3]
	fmt.Println(featuredPrice)

	// concepts
	// arrays are stored in memory, kind of like pointers (but different concept)
	// slices are windows into specific parts of the stored data
	// so if you modify a value from that slice, you are indeed modifying the original array
	// so when creating a slice you are NOT copying it, but still accessing that one array you have in memory
	//
	// some quirks
	// slices have length and capacity properties
	fmt.Println(len(featuredPrice), cap(featuredPrice))
	// length is number of items in slice or array
	// capacity is how many other elements you can access
	// why is it not 4, because you can always access more from right
	// not left... i think i need to deep dive this to understand more
	// playing around with this to understand
	sliceA := price[:1]
	fmt.Println(sliceA)
	// return one item
	sliceA = sliceA[:3]
	fmt.Println(sliceA)
	// teturns 3
	fmt.Println(len(sliceA), cap(sliceA))

	// go arrays are not dynamic, the size must be set
	// this creates a dynamic array in the background in memory
	// even though it is a slice. go automatically ditches the array
	// if it grows beyond the bounds of the original array and createa new one behidn the scenes
	values := []float64{22.00}

	// with the above, you still can't access values that don't exist or set them to an index that dne
	// however go gives us some function
	// this creates a NEW array though (and returns it)
	// i.e., original slice does not change
	// as it is a new array in memory
	newValues := append(values, 1.11)
	fmt.Println(newValues)
	// can of course overwrite that old slice in memory
	values = append(values, 1.11)
	fmt.Println(values)

	// merging slices
	altValues := []float64{22.22, 100.2}
	values = append(values, altValues...)
	// this  `...` is unpacking the vlaues so you aren't adding the slice, but rather the comma sep values
	fmt.Println(values)

	// make function
	// this tells go to make an array with a length of 2 in memory.
	// this the other way []string{} makes an array of lenght 0
	// then we replace it with new array with the appends.
	userName := make([]string, 2, 5)
	// this is useful if you want to assign specific values
	// to specific indexes on an empty array
	// so the `2` tells go to maek an array with 2 length
	// but the 5 says allocate enough memory so that it could
	// have a capacity of 5, so go allocates enough memory

	userName = append(userName, "Isaac")
	userName = append(userName, "Ana")
	fmt.Println(userName)

	// THIS ALSO WORKS FOR MAPS
	// but the argument is to give the intended length of that map
	// so that memory is allocated

	examplemaps.MapsExample()
}
