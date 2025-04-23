// Escriba un programa que reciba como argumento un nombre de archivo y muestre por
// consola su contenido. Si el archivo no existe se debe mostrar al usuario un mensaje y
// terminar.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) > 1 {

		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Error al abrir el archivo")
			return
		}
		// Utilizamos esto para cerrar el archivo una vez que se termine de leerlo
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error al leer el archivo")
			return
		}

		fmt.Printf("Contenido del archivo:\n")
		fmt.Printf(string(content))

	} else {
		fmt.Printf("No se ingresó ningun argumento")
	}

}
