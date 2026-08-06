package main

import (
	"errors"
	"fmt"
	"time"

	// this is importing a custom struct pacakge
	"example.com/structs/boardgame"
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
// as written, the receiver argument works like other arguments in that
// they make a copy in memory, so if we made a clear app user func
// doing someting like u.firtName = "", it would be setting the copy `u`, not
// the value stored in the struct, see mutating below
func (u user) outputUserDetailsMethod() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// as stated above, to mutate you should use the pointer instead so the
// original instance is mutated
func (u *user) clearUserFirstName() {
	u.firstName = ""
}

// constructors
// this isn't necessarily a feature baked into go
// like classes in other languages, but rather a convention
// you typically see in go code in which the job of this (constructor)
// function is to to create such a struct
// convention is to name it `newStructName`
// you can do the method we did below using struct literal notation still,
// however, this is common convention
// also note that you can use a pointer for a constructor because this is
// really making a copy, however, for this amount of data, it is not necessary
// see newUserWithPointer below
func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

// constructor with pointer and validation
func newUserWithPointer(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Can't be blank")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
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

	// using the constructor
	appUserFromCon := newUser(userFirstName, userLastName, userBirthdate)
	fmt.Println("App User from Constructor", appUserFromCon)

	// using constructor with pointer + validation
	var appUserFromConPointer *user
	appUserFromConPointer, err := newUserWithPointer(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("App user form constructor with pointer:", appUserFromConPointer)

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
	appUser.clearUserFirstName()
	appUser.outputUserDetailsMethod()

	// this is an example of splitting the struct in its own package
	var game *boardgame.Boardgame
	game, _ = boardgame.NewBoardgame("Dominion", 100)
	fmt.Println(game)
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
	// scanln ends the scan when enter is pressed
	fmt.Scanln(&value)

	return value
}
