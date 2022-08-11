package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	//TODO: Implementar forma de esperar uma função terminar a execucao
	wg.Add(10)
	for i := 0; i < 5; i++ {

		go func() {
			defer wg.Done()
			DemorarParaExecutar(2 * time.Second)
		}()
	}

	wg.Wait()

	//TODO: E se eu registrar registradores a menos no waitgroup?
	//TODO: E se eu registrar valores a mais no waitgroup?
}

func DemorarParaExecutar(t time.Duration) {
	fmt.Println("Começando a execução do método")

	time.Sleep(t)

	fmt.Println("Finalizando a execução do método")
}
