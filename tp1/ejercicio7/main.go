// Simular el acceso concurrente a un log compartido donde 10 goroutines, cada una
// representando un nodo, intentan registrar un evento crítico cada 0.5 segundos (por ejemplo:
// "nodo-3: temperatura alta" o "nodo-7: pérdida de conexión"); cada evento debe escribirse en
// un slice de strings protegido por sync.Mutex para garantizar la integridad del log
package main

import (
	"fmt"
	"sync"
	"time"
)

var sliceCompartido []string
var mutex sync.Mutex

func nodo1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 1: Log generado por nodo 1")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 2: Log generado por nodo 2")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo3(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 3: Log generado por nodo 3")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo4(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 4: Log generado por nodo 4")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo5(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 5: Log generado por nodo 5")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo6(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 6: Log generado por nodo 6")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo7(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 7: Log generado por nodo 7")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo8(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 8: Log generado por nodo 8")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo9(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 9: Log generado por nodo 9")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func nodo10(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		sliceCompartido = append(sliceCompartido, "nodo 10: Log generado por nodo 10")
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	go nodo1(&wg)
	go nodo2(&wg)
	go nodo3(&wg)
	go nodo4(&wg)
	go nodo5(&wg)
	go nodo6(&wg)
	go nodo7(&wg)
	go nodo8(&wg)
	go nodo9(&wg)
	go nodo10(&wg)

	wg.Wait()

	fmt.Println("Contenido del slice: ")
	for _, mensaje := range sliceCompartido {
		println(mensaje)
	}

}
