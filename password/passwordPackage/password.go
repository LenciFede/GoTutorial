package password

import "regexp"

func ValidatePassword(password string) bool {
	contieneNumero, _ := regexp.MatchString(".*\\d.*", password)
	contieneMayuscula, _ := regexp.MatchString(".*[A-Z].*", password)
	contieneAlMenos8Caracteres := len(password) >= 8

	if contieneNumero && contieneMayuscula && contieneAlMenos8Caracteres {
		return true
	} else {
		return false
	}

}
