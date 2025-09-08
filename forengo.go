package main

import "fmt"

func main() {
	numeros :=
		[]int{10, 20, 30}
	palabra := "Go!"

	for i := 0; i < 5; i++ {
		fmt.Println("iteracion", i)

	}

	for i, v := range numeros {
		fmt.Println("indice: %d, valor: %d", i, v)

	}

	for _, v := range numeros {

		fmt.Println("valor", v)

	}

	for i, r := range palabra {
		fmt.Printf("posición: %d, Runa: %c", i, r)
	}

}
