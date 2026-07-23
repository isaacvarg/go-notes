package main

import "fmt"

// for loops
// for i := 0; i < 2; i++
// similar to other langs
// set a variable, condition to keep going, what happens after each run
//
// there is also a variation of for loop such that
// for someCondition { do someting}
// where someCondition returns boolean
// the loop will continue while true
//
// there is a switch statement
// that functions very similarly to Ts
// don't need break after each case
// break for default does exit the switch though
//

func main() {
	currentNumber := 10

	fmt.Println("Welcome to Number Changes!")
	// this is simply an infinite loop
	for {
		fmt.Println("")
		fmt.Println("What would you like to do?")
		fmt.Println("1. Look at the number")
		fmt.Println("2. Add to the number")
		fmt.Println("3. Remove from the number")
		fmt.Println("4. Exit")

		fmt.Println("")
		fmt.Print("What is your choice?: ")

		var choice int
		fmt.Scan(&choice)

		// so we could do multiple if statements, but
		// if we use else if then the others are not ran
		// if the first is true, same as other languages
		// i am sure there is a switch control through similar to ts
		if choice == 1 {
			fmt.Println("Current number is", currentNumber)
		} else if choice == 2 {
			fmt.Print("How much to add? ")
			var addAmount int
			fmt.Scan(&addAmount)

			if addAmount <= 0 {
				fmt.Println("Must be great than zero")
				// this stops other code from being executed, like other langs
				// return

				// continue is keyword that exits the current loop and
				// continues on to the next iteration without
				// exiting the app like return does
				continue
			}

			// shorthand for currentNumber + addAmount
			currentNumber += addAmount

			fmt.Println("New number is", currentNumber)
		} else if choice == 3 {
			fmt.Print("How much to remove? ")
			var subAmount int

			fmt.Scan(&subAmount)

			if subAmount <= 0 {
				fmt.Println("Must be greater than zero")
				continue
			}

			// now running this if regardless of previous if
			// really could have used && operator but
			if subAmount > currentNumber {
				fmt.Println("Woah there, we subtract more than we have")
				continue
			}

			// shorthand for currentNumber + addAmount
			currentNumber -= subAmount

			fmt.Println("New number is", currentNumber)
		} else {
			fmt.Println("bye bye!")
			// could use return here but then the
			// thanks for playing is not reachable
			// break can only used to exit loops
			break
		}
	}

	fmt.Println("Thanks for playing!")
}
