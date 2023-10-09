package main

import (
	"fmt"

	correo "github.com/LenciFede/GoTutorial/correo/correoPackage"
)

func main() {
	fmt.Printf("---------------------\n")
	correo.EnviarMailSoloDestinatario("Fede")
	fmt.Printf("---------------------")
	fmt.Printf("\n")
	fmt.Printf("---------------------")
	fmt.Printf("\n")
	correo.EnviarMailSinMensaje("Fede", "TDD")
	fmt.Printf("---------------------")
	fmt.Printf("\n")
	fmt.Printf("---------------------")
	fmt.Printf("\n")
	correo.EnviarMailCompleto("Fede", "TDD", "Esto es un test")

	fmt.Printf("---------------------\n")

}
