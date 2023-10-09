package main

import (
	"fmt"

	password "github.com/LenciFede/GoTutorial/password/passwordPackage"
)

func main() {
	var esContraseñaValida = password.ValidatePassword("Hola1234")
	fmt.Println("Es una contrasela valida? ", esContraseñaValida)
}
