// Escriba un programa que mediante el uso de un mutex global, escriba dos funciones donde:
// la función a() debe bloquear el mutex, invocar a la función b() y desbloquear el mutex; la
// función b() debe bloquear el mutex, imprimir “Hola mundo” y desbloquear el mutex. La
// gorutina principal debe invocar a la función a(). Explica que sucede al ejecutar el programa.

package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func a() {
	mutex.Lock()
	b()
	mutex.Unlock()
}

func b() {
	mutex.Lock()
	fmt.Println("Hola mundo")
	mutex.Unlock()
}

func main() {

	go a()

}

// Lo que sucede es que al invocar la funcion a el mutex queda bloqueado antes de invocar a b
// y esto hace que el programa no pueda acceder a b ya que necesitamos bloquear el mutex nuevamente
// para imprimir el hola mundo y ya se encuentra bloqueado en la ejecucion de la funcion a
