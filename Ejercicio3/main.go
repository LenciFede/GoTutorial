package main

func sumar(numero1 int, numero2 int) int {
	var resultado = numero1 + numero2
	return resultado
}

func restar(numero1 int, numero2 int) int {
	var resultado = numero2 - numero1
	return resultado
}

func operacionSumaYResta(numero1 int, numero2 int) (suma int, resta int) {
	suma = sumar(numero1, numero2)
	resta = restar(numero2, numero1)
	return suma, resta
}

func main() {

}
