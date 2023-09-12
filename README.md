# greetings
Sección 7 curso de golang

Este paquete proporciona una forma simple de obtener saludos personalizados

## Instalación
Ejecuta el siguiente comando para instalar el paquete

```bash
go get -u github.com/LeandroMauro/greetings
```

#las 3 comillas y la palabra bash, hacen que la linea del go get se vea como una porción de código

## Uso

Aqui tienes un ejemplo de como utilizar el paquete en tu código....

```go
package main

import (
	"fmt"

	"github.com/LeandroMauro/greetings"
)

func main() {
	message, err := greetings.Hello("Alex")

	if err != nil {
		fmt.Println("Ocurrió un error: ", err)
		return
	}

	fmt.Println(message)
}
```