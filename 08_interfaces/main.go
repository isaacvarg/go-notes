package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

// i know functionally this is different in goland to typescript
// but the developer experience is kind of like passing generics in
// typescript... except maybe stricter... i love it

// so interfaces basically describe what should exist for the method
// convention is if the interface only has one method that it should be the
// method name with er
type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type outputtable interface {
	Saver
	Displayer
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	// now using the func that uses the interface
	// so this works because go is looking at the note struct or package and ensuring that is
	// indeed has a method called Save() which it does
	// you don't need to link it up it just works
	err = saveData(userNote)
	if err != nil {
		return
	}
}

// this is saying that the data will be of some type that
// adheres to that interface
// this is k,idn of
func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Not saved.")
		return err
	}

	fmt.Println("Saving successful")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

// if simply used the displayer interface as the data type this would give us an error
// because that was go realized that data is not gauranteed to have a Save method
// so we use a embedded interface
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// ANY value
// instead if `interface{}` you can use `any` keyword
// obvious this is like in typescript, you should really only use this in
// very specific scenarios
func printSomething(value interface{}) {
	fmt.Println(value)
}

// type switch
// go provides a special switch statement for types
func printSomethingSwitch(value any) {
	// this is also a cool syntax for getting type data from a value
	// this is checking if the value is an int
	// now typedValue will be an int after this check, but will obv lead
	// to some errors downstream
	typedValue, isTypeWeWant := value.(int)
	fmt.Println(typedValue, isTypeWeWant)

	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("float64:", value)
		// does nothing for one that isn't int or float64
		// because we don't have default
		// the value would be accepted but nothing would happpen
	}
}

// GENERICS
// this is what i was fishing for in some practice projects I was doing
// this really helps when you have a function that you want to accept multiple types
// e.g., a float or int without accepting any kind of value then type checking with `any`
//
// the brackets turn this into a generic and inside the brackets you add a placeholder name
// convention is T, kind of like typescript
// then you tell the types allowed as T after separated by |
// so this accepts a and b when they are type T, which is float64 or int, and returns T
// I like this syntax more than generics in typescript
func add[T float64 | int](a, b T) T {
	return a + b
}
