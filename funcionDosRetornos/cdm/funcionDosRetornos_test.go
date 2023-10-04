package funcionDosRetornos_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	funcionDosRetornos "github.com/LenciFede/GoTutorial/funcionDosRetornos/cdm"
)

func TestOperacionSumaYResta(t *testing.T) {
	var suma, resta = funcionDosRetornos.OperacionSumaYResta(1, 1)

	assert.Equal(t, suma, 2)
	assert.Equal(t, resta, 0)
}
