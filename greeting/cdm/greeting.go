package greeting

func Greeting(name string) string {
	var greet = ""
	if name == "" {
		greet = "Hola Extraño"
		return greet
	} else {
		greet = "Hola " + name
		return greet
	}
}
