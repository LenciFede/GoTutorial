/*
Programa que simule el envío de un mail imprimiendo por pantalla el mensaje

enviando correo...
destinatario: <destinatario>
asunto: <asunto>
mensaje: <mensaje>

Para los distintos casos que se van a plantear a continuación la idea es reutilizar la
función de “simulación de envio de mail”

# Escribir el código dentro de un paquete llamado correo y llamar desde el main

a. Crear función no exportada llamada simulaEnvio que reciba por parámetros un
destinatario, asunto y mensaje del tipo string e imprima por pantalla “enviando
correo...” y debajo el destinatario, el asunto y el mensaje. En el caso de que el
parámetro sea vacío imprimir un guión “-”

b. Función exportada que me permita enviar un mail con un destinatario (sin asunto	y sin mensaje)

c. Función exportada que me permita enviar un mail con un destinatario y asunto
(sin mensaje)

d. Función exportada que me permita enviar un mail con un destinatario, asunto y
mensaje
*/
package correo_test

import (
	"testing"

	correo "github.com/LenciFede/GoTutorial/correo/correoPackage"
)

func dadoUnDestinatarioUnAsuntoYUnMensaje(destinatario string, asunto string, mensaje string) (string, string, string) {
	dest := destinatario
	asun := asunto
	mens := mensaje

	return dest, asun, mens
}

func queCreaUnMailParaPorPantalla(destinatario string, asunto string, mensaje string) {
	correo.EnviarMailCompleto(destinatario, asunto, mensaje)
}

func TestCreaUnMailConDestinatarioMensajeYAsunto(t *testing.T) {
	var destinatario, mensaje, asunto = dadoUnDestinatarioUnAsuntoYUnMensaje("Fede", "PruebaTDD", "Esto es una prueba TDD")
	queCreaUnMailParaPorPantalla(destinatario, mensaje, asunto)
}
