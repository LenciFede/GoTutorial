package correo

import "fmt"

func simulaEnvio(destinatario string, asunto string, mensaje string) {
	if asunto == "" {
		asunto = "-"
	}
	if mensaje == "" {
		mensaje = "-"
	}
	var mail = fmt.Sprintf("Enviando correo...\nDestinatario: %s\nAsunto: %s\nMensaje: %s\n", destinatario, asunto, mensaje)
	fmt.Printf(mail)
}

func EnviarMailSoloDestinatario(destinatario string) {
	simulaEnvio(destinatario, "", "")
}

func EnviarMailSinMensaje(destinatario string, asunto string) {
	simulaEnvio(destinatario, asunto, "")
}

func EnviarMailCompleto(destinatario string, asunto string, mensaje string) {
	simulaEnvio(destinatario, asunto, mensaje)
}
