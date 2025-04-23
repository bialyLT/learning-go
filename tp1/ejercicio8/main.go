// Simular un sistema de monitoreo donde 1 goroutina, envían un ping a una lista de 3 nodos
// (nodo-1, nodo-2, nodo-3) cada 2 segundos; cada ping simula una latencia aleatoria entre 100
// y 500 milisegundos, la gorutina debe guardar el nodo con menor latencia de respuesta de
// cada ronda en un slice. Luego de 10 rondas se debe imprimir los resultados y terminar.
// Consulta: se debe proteger el slice con mutex.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func nodo1(ch chan int) {
	ping := rand.Intn(401) + 100
	ch <- ping
}

func nodo2(ch chan int) {
	ping := rand.Intn(401) + 100
	ch <- ping
}

func nodo3(ch chan int) {
	ping := rand.Intn(401) + 100
	ch <- ping
}

func recibirPings(canal1, canal2, canal3 chan int, nodos *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		go nodo1(canal1)
		go nodo2(canal2)
		go nodo3(canal3)

		select {
		case ping := <-canal1:
			fmt.Printf("Ping from nodo 1: %d ms\n", ping)
			*nodos = append(*nodos, "nodo 1")
			time.Sleep(2 * time.Second)
		case ping := <-canal2:
			fmt.Printf("Ping from nodo 2: %d ms\n", ping)
			*nodos = append(*nodos, "nodo 2")
			time.Sleep(2 * time.Second)
		case ping := <-canal3:
			fmt.Printf("Ping from nodo 3: %d ms\n", ping)
			*nodos = append(*nodos, "nodo 3")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {

	var wg sync.WaitGroup

	canal1 := make(chan int, 1)
	canal2 := make(chan int, 1)
	canal3 := make(chan int, 1)

	var nodos_mas_rapidos []string

	wg.Add(1)

	go recibirPings(canal1, canal2, canal3, &nodos_mas_rapidos, &wg)

	wg.Wait()

	fmt.Println("Los nodos mas rapidos fueron: ")
	for i := range nodos_mas_rapidos {
		fmt.Println(nodos_mas_rapidos[i])
	}
}
