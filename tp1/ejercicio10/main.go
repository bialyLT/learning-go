// Desarrolle un programa que tenga una variable global x iniciada en 0 (cero) y una función
// incrementar() que incremente en 5 a la variable x. La gorutina principal debe lanzar 100
// gorutinas que invoquen a la función incrementar() y luego imprimir el valor de x. Ejecute su
// programa usando la bandera -race para detectar si hay una carrera de datos. Además, el valor
// final de x debe ser 500, pero es posible que observe que a veces es 490 o 495 u otros
// valores. Usando WaitGroup y Mutexes, corrija su programa para que imprima el valor
// correcto y no tenga una carrera de datos

package main

import (
	"fmt"
	"sync"
)

var x = 0
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	x += 5
	mutex.Unlock()
}

func principal(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		go incrementar(wg)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(101)

	go principal(&wg)

	wg.Wait()

	fmt.Println(x)

}
