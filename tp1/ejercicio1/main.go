// Escriba una función SumarPares que reciba un slice de enteros y devuelva la suma de los
// números pares. Implemente un programa que demuestre su funcionamiento.

package main

import "fmt"

func SumarPares(numeros []int) int {
	suma := 0
	for _, n := range numeros {
		if n%2 == 0 {
			suma += n
		}
	}
	return suma
}
func main() {
	slicesDeNumeros := []int{1, 2, 3, 5, 6, 3}
	fmt.Println(SumarPares(slicesDeNumeros))
}
