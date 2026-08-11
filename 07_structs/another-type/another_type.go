package main

// types can also be used to expand existing types
// this is a stupid example but you will get the idea
// so here we are taking string and extending it by assigning it an alias

type str string

// now we are assigning a new method to string by passing a receiver func
func (text str) log() {
 fmt.Println(text)
}

// now you can access a method on string that previous didn't exist
func AnotherType() {
	var myString str = "Isaac"

	myString.log()
}
