package main

import "fmt"

type Usuario struct {
	Nombre string
	edad   int
}

type Admin struct {
	Usuario
	Nivel int
}

type Hablador interface {
	Hablar() string
}

func (u Usuario) Hablar() string {
	return "Hola!"
}

func anunciar(h Hablador) {
	fmt.Println(h.Hablar())
}

func (u *Usuario) saludar() string {
	return fmt.Sprintf("hola, soy %s y tengo %d", u.Nombre, u.edad)
}

func newUsuario(nombre string, edad int) *Usuario {
	return &Usuario{Nombre: nombre, edad: edad}
}

func main() {

	u := &Usuario{Nombre: "Rafa", edad: 34}
	fmt.Println(u.saludar())
	anunciar(u)
	a := Admin{Usuario: *u, Nivel: 12}
	fmt.Println(a.Usuario.Nombre)
}
