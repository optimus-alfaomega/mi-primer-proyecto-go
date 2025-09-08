package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	fmt.Println("hola", nombre)
}

func main() {

	go saludar("Rafael")
	go saludar("Nequi")
	time.Sleep(100 * time.Millisecond)

}
