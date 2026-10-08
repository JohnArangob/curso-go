package main

import (
	"fmt"
)

type AnimalVolador interface {
	Volar
	Correr
	Carnivoro
}
type AnimalCazador interface {
	Correr
	Carnivoro
	Cazar
}
type Animal interface {
	Correr
	Carnivoro
}
type Volar interface {
	Volar() string
}

type Correr interface {
	Correr() string
}
type Cazar interface {
	Depreda() string
}
type Carnivoro interface {
	Comer() string
}

type Pajaro struct {
	Nombre string
}

type Perro struct {
	Nombre string
}

type Leon struct {
	Nombre string
}

func (p Pajaro) Volar() string {
	return p.Nombre + " vuela."
}

func (p Pajaro) Correr() string {
	return p.Nombre + " corre."
}

func (p Pajaro) Comer() string {
	return p.Nombre + " come insectos."
}

func (p Perro) Correr() string {
	return p.Nombre + " corre."
}

func (p Perro) Comer() string {
	return p.Nombre + " come carne."
}

func (l Leon) Correr() string {
	return l.Nombre + "Corre mucho."
}

func (l Leon) Comer() string {
	return l.Nombre + "Come carne."
}

func (l Leon) Depreda() string {
	return l.Nombre + "Saca sus garras para cazar."
}

func main() {
	fmt.Println("¿Qué animal quieres consultar?")
	fmt.Println("1. Pájaro (Correcaminos)")
	fmt.Println("2. Perro (Max)")
	fmt.Println("3. Leon (Rayitas)")
	fmt.Print("Elige 1, 2 o 3: ")

	var opcion int
	if _, err := fmt.Scan(&opcion); err != nil {
		fmt.Println("Entrada no válida; escribe 1, 2 o 3.")
		return
	}

	switch opcion {
	case 1:
		var animal AnimalVolador = Pajaro{Nombre: "Correcaminos"}
		fmt.Println("Tipo: pájaro")
		fmt.Println(animal.Correr())
		fmt.Println(animal.Comer())
		fmt.Println(animal.Volar())
	case 2:
		var animal Animal = Perro{Nombre: "Max"}
		fmt.Println("Tipo: perro")
		fmt.Println(animal.Correr())
		fmt.Println(animal.Comer())
	case 3:
		var animal AnimalCazador = Leon{Nombre: "Rayitas "}
		fmt.Println(animal.Correr())
		fmt.Println(animal.Comer())
		fmt.Println(animal.Depreda())
	default:
		fmt.Println("Opción no válida; elige 1 o 2.")
	}
}
