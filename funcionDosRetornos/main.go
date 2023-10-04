package main

import (
	"fmt"

	funcionDosRetornos "github.com/LenciFede/GoTutorial/funcionDosRetornos/cdm"
)

func main() {
	var suma, resta = funcionDosRetornos.OperacionSumaYResta(2, 2)
	var salida = fmt.Sprintf("Valor de la suma: %d,\nValor de la resta: %d", suma, resta)
	fmt.Println(salida)

	var soloSuma, _ = funcionDosRetornos.OperacionSumaYResta(2, 2)
	var nuevaSalida = fmt.Sprintf("Valor de la suma: %d,\n", soloSuma)
	fmt.Println(nuevaSalida)

}
