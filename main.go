package main

import "fmt"

type Animal interface {
	hablar() string
}

type Perro struct{}
type Gato struct{}
type Vaca struct{}

func (Perro) hablar() string { return "Guau!" }
func (Gato) hablar() string  { return "Miau!" }
func (Vaca) hablar() string  { return "Muuu!" }

func main() {
	mascotas := []Animal{Perro{}, Vaca{}}
	fmt.Println(mascotas[1].hablar())

}
