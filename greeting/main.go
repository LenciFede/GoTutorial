package main

import (
	"fmt"

	greeting "github.com/LenciFede/GoTutorial/greeting/cdm"
)

func main() {
	var greetWithName string = greeting.Greeting("Fede")
	fmt.Printf(greetWithName)
	fmt.Printf("\n")
	var greetWithoutName = greeting.Greeting("")
	fmt.Printf(greetWithoutName)
	fmt.Printf("\n")
}
