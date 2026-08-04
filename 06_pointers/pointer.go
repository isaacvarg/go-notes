package main

import "fmt"

// other notes
// all values in go have an nil value
// for int it is 0, float it is 0.0, string is "", etc
// for a pointer, nil represents
// absense of a value address
// i.e., a pointer pointing at no address / no value in memory
//
// You absolutely should not try and not duplicate the value as shown here
// this is just to exemplify how pointers work... these values are incredibly small
// and okay to garbage collect... this really only becomes an issue in specific use cases
//

func main() {
	// so this is creating a value of 33 at a certain address
	age := 33

	// the & creates a pointer to the address where
	// the value is stored
	// this creates a type of *int, where
	// * indicates its a pointer
	// this could be done in two steps
	agePointer := &age

	// shows us actual address
	fmt.Println(agePointer)

	// to get the actual value at the address you must
	// dereference by adding an * infront of the pointer
	fmt.Println("Value from pointer", *agePointer)

	// using pointer as the argument
	// could also use &age
	fmt.Println("pointer function value", adultYearsPointer(agePointer))

	// then these are technically creating new copies
	// of the value at new addresses
	// since they are being passed as arguments
	fmt.Println(age)
	fmt.Println(adultYears(age))

	// this should read 33 before we mutate the value
	fmt.Println(age)
	// this is an example of mutation with pointers
	setAgeToAdultYears(&age)
	// this should read 15 since we mutated it
	fmt.Println(age)
}

func adultYears(age int) int {
	return age - 18
}

func adultYearsPointer(age *int) int {
	// so age - 18 is err
	// because there is a type mismatch
	// also you cannot do arithm. on pointer
	// so you must dereference before use
	return *age - 18
}

func setAgeToAdultYears(age *int) {
	// notice this doesn't return anything
	// and simply mutates the value by accessing it
	// from the memory address
	*age = *age - 18
}
