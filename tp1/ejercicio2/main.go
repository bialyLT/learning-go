// Desarrolle un programa que lea una línea de texto desde la entrada estándar y cuente e
// imprima cuántas palabras tiene. Busque ayuda en los paquetes strings y fmt.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func calcularCantidadPalabras(texto string) int {

	cantidadPalabras := 0

	for i := range len(texto) {
		if texto[i] != ' ' && i != len(texto)-1 {
			continue
		}
		if texto[i-1] != ' ' {
			cantidadPalabras++
		}
	}

	return cantidadPalabras
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un texto:")
	if scanner.Scan() {
		textoIngresado := scanner.Text()
		fmt.Println("La cantidad de palabras es de: ", calcularCantidadPalabras(textoIngresado))
	}
}
