// Simular un middleware donde un único publicador envía cada 1 segundo un mensaje (por
// 	ejemplo: "evento-X") a 3 suscriptores. Cada suscriptor está representado por una goroutine
// 	que escucha su propio canal y muestra los eventos recibidos. El sistema debe permitir que
// 	todos los suscriptores reciban el mismo mensaje simultáneamente y en igual orden

package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func suscriptor(ch <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for mensaje := range ch {
		fmt.Println("suscriptor N", id, "recibio el evento: ", mensaje)
	}
}

func publicador(ch1, ch2, ch3 chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		evento := fmt.Sprintf("evento %d", i+1)
		ch1 <- evento
		ch2 <- evento
		ch3 <- evento
		time.Sleep(1 * time.Millisecond)
	}
	close(ch1)
	close(ch2)
	close(ch3)
}

func main() {

	var wg sync.WaitGroup

	canal1 := make(chan string)
	canal2 := make(chan string)
	canal3 := make(chan string)

	wg.Add(4)

	go publicador(canal1, canal2, canal3, &wg)

	go suscriptor(canal1, 1, &wg)
	go suscriptor(canal2, 2, &wg)
	go suscriptor(canal3, 3, &wg)

	wg.Wait()

}
