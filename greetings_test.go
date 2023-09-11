package greetings

import (
	"regexp"
	"testing"
)

// Para probar modulos de go
// la clase debe terminar con "_test"
// Y la función comenzar con "Test"

// Para ejecutar las pruebas en la consola
// go test o sino go test -v
func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`) //Esta comilla se llama backtick o acento grave
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
