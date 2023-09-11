package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// CREACION DE UN MODULO
// Hello devuelve un saludo para la persona específica
// PD: Las funciones "globales" o a ser utilizadas por otro módulo deben comenzar si o si con mayúscula
func Hello(name string) (string, error) {
	//Se agrega manejo de errores
	if name == "" {
		return name, errors.New("Nombre vacío")
	}
	//Devuelve un saludo que incluye el nombre en el mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos devuelve un mapa que asociada cada nombre de las personas
// Con un mensaje de saludo/greeting
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludo
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
