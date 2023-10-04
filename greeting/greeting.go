package greeting

import "fmt"

func Greeting(name string) string {
	var greet = ""
	if name == "" {
		greet = "Hola Extraño"
		return greet
	} else {
		greet = "Hola " + name
		return greet
	}
}

func main() {
	var greetWithName string = Greeting("Fede")
	fmt.Printf(greetWithName)
	fmt.Printf("\n")
	var greetWithoutName = Greeting("")
	fmt.Printf(greetWithoutName)
	fmt.Printf("\n")
}
