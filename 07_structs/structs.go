package main

import (
	"fmt"
	"time"
)

// so far as I can tell this is similar to type in Typescript
// we can name this in capital case to export
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// METHODS
// so the (u user) is called a receiver arguemnt (or just receiver)
// the `u` allows one to access the values of the struct
// and the `user` adds this func to the user struct
// so now we don't need to pass the arguments because this method
// has access to the data of the struct it is apart of via the `u`
func (u user) outputUserDetailsMethod() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// this is an instance of a struct
	// this is also called struct literal  or composite literal
	var appUser user

	// to me this is functioning like a ts object
	// you define it as an instance by the structName{}, no space is convention
	// trailing comma is required
	// this is struct literal notation
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// if you are adding the values in the order of the struct
	// you can ommit the keyword
	// my neovim actually displays the keyword nicely on to the left when omitting it
	// e.g.,
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }
	//
	// you can also create the nul value of the struct by:
	// appUser = user{}
	//
	// values can also be ommitted by not
	// placing then in the struct literal
	// and that value in would be the nul value for that type
	// e.g., ommitting birthdate would mean the value is ""

	// this is passing struct as a parameter
	outputUserDetails(appUser)
	// this is passing pointer of the struct As a paramater
	altOutputUserDetails(&appUser)

	// now using the method instead, we don't need to pass argunments
	appUser.outputUserDetailsMethod()
}

// now let's actually use the struct
// as a parameter
// appUser is the parameter name, user is the value type
// as we see here, it uses dot notation like methods and objects
func outputUserDetails(appUser user) {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

// this is another example of using the pointer instead of
// making a copy of the var, however, with this size of data
// not a big deal
func altOutputUserDetails(appUser *user) {
	// you would think you would need to dereference the pointer
	// e.g., (*appUser.firstName)
	// although technically correct way to access a field on a struct pointer
	// go allows you to access it with this notation without derefencing
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println(err)
	}
	return value
}
