/*
Crear un método que permita validar la seguridad de una contraseña retornando un
error si no cumple con algún criterio especificado. La contraseña debe tener al menos 8
caracteres, un número y una mayúscula.

TIP: usar regex para simplificar la detección de las reestricciones ;)
regexp.MatchString(".*\\d.*", password)
regexp.MatchString(".*[A-Z].*", password)

a.Test que falle si tiene menos de 8 caracteres
b.Test que falle si no tiene un número
c.Test que falle si no tiene una mayúscula
d.Test que pase la validación y que la contraseña tenga más de un número
*/

package password_test

import (
	"testing"

	password "github.com/LenciFede/GoTutorial/password/passwordPackage"
	"github.com/stretchr/testify/assert"
)

func TestPasswordMenorOchoCaracteres(t *testing.T) {
	var contraseñaValida = password.ValidatePassword("Hola1234")
	assert.Equal(t, contraseñaValida, false)
}

func TestPasswordSinNumeros(t *testing.T) {
	var contraseñaValida = password.ValidatePassword("Holahola")
	assert.Equal(t, contraseñaValida, false)
}

func TestPasswordSinMayuscula(t *testing.T) {
	var contraseñaValida = password.ValidatePassword("hola1234")
	assert.Equal(t, contraseñaValida, false)
}

func TestPasswordCorrecta(t *testing.T) {
	var contraseñaValida = password.ValidatePassword("hola1234")
	assert.Equal(t, contraseñaValida, false)
}
