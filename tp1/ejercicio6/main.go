// Diseñar un sistema en anillo donde cinco nodos, representados por goroutines, se envían
// mensajes de heartbeat cada 1 segundo entre sí de manera cíclica a través de canales. El
// sistema debe funcionar por 1 minuto y terminar.
package main

import (
	"fmt"
	"sync"
	"time"
)

func nodo1(entrada, salida chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		mensaje := <-entrada
		salida <- "Hola desde el nodo 1"
		time.Sleep(1 * time.Second)
		fmt.Println(mensaje)
	}
}

func nodo2(entrada, salida chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		mensaje := <-entrada
		salida <- "Hola desde el nodo 2"
		time.Sleep(1 * time.Second)
		fmt.Println(mensaje)
	}
}

func nodo3(entrada, salida chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		mensaje := <-entrada
		salida <- "Hola desde el nodo 3"
		time.Sleep(1 * time.Second)
		fmt.Println(mensaje)
	}
}

func nodo4(entrada, salida chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		mensaje := <-entrada
		salida <- "Hola desde el nodo 4"
		time.Sleep(1 * time.Second)
		fmt.Println(mensaje)
	}
}

func nodo5(entrada, salida chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		mensaje := <-entrada
		salida <- "Hola desde el nodo 5"
		time.Sleep(1 * time.Second)
		fmt.Println(mensaje)
	}
}

func main() {
	var wg sync.WaitGroup

	// declaramos los canales agregando un buffer de tamaño 1 para lograr que
	// se envien los mensajes aunque el receptor no este listo y no se bloquee al principio
	canal1_2 := make(chan string, 1)
	canal2_3 := make(chan string, 1)
	canal3_4 := make(chan string, 1)
	canal4_5 := make(chan string, 1)
	canal5_1 := make(chan string, 1)

	wg.Add(5)

	go nodo1(canal5_1, canal1_2, &wg)
	go nodo2(canal1_2, canal2_3, &wg)
	go nodo3(canal2_3, canal3_4, &wg)
	go nodo4(canal3_4, canal4_5, &wg)
	go nodo5(canal4_5, canal5_1, &wg)

	canal5_1 <- "Inicio del ciclo"

	wg.Wait()
}
