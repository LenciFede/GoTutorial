package greeting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"GoTutorial/greeting"
)

// TestGreeting usa el paquete que viene para testing en golang.
func TestGreeting(t *testing.T) {
	var completeGreeting string = greeting.Greeting("Fede")

	if completeGreeting != "Hola Fede" {
		t.Error("El mensaje deberia mostrar el nombre")
	}
}

// TestGreetingConTestify usa el paquete assert de testify para probar,
// esta dependencia la manejamos con go modules.
func TestGreetingConTestify(t *testing.T) {
	var completeGreeting string = greeting.Greeting("Fede")
	assert.Equal(t, completeGreeting, "Hola Fede")
}
