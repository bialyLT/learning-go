// Escriba una función que convierta de °C a °F y otra de °F a °C. Luego realice un menú para
// elegir qué conversión hacer y pida los datos por teclado
package main

import "fmt"

func menu() {

	var seleccion int
	var fahrenheit, celsius float64

	fmt.Println("Elija una opción: ")
	fmt.Println("1. fahrenheit a celsius")
	fmt.Println("2. celsius a Fahrenheit")
	fmt.Println("3. Terminar el programa")

	fmt.Scanln(&seleccion)

	switch seleccion {
	case 1:
		fmt.Println("ingrese los grados Fahrenheit: ")
		fmt.Scanln(&fahrenheit)
		fmt.Printf("%v grados fahrenheit son %v grados celsius\n", fahrenheit, convertir(1, celsius, fahrenheit))
		volverAlMenu()

	case 2:
		fmt.Println("ingrese los grados Celsius: ")
		fmt.Scanln(&celsius)
		fmt.Printf("%v grados celsius son %v grados fahrenheit\n", celsius, convertir(2, celsius, fahrenheit))
		volverAlMenu()
	case 3:
		break
	default:
		fmt.Println("haz seleccionado una opcion incorrecta")
		menu()

	}
}

func convertir(tipoDeConversion int, celsius, fahrenheit float64) float64 {
	if tipoDeConversion == 1 {
		return (fahrenheit - 32) * 5 / 9
	} else if tipoDeConversion == 2 {
		return (celsius * 9 / 5) + 32
	}
	return 0
}

func volverAlMenu() {
	var seleccion int
	fmt.Println("1. Volver al menu")
	fmt.Println("2. Terminar con el programa")
	fmt.Scanln(&seleccion)
	switch seleccion {
	case 1:
		menu()
	case 2:
		break
	default:
		fmt.Println("haz seleccionado una opcion incorrecta")
		volverAlMenu()
	}

}

func main() {
	menu()
}
