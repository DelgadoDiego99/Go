package main

import "fmt"

func main() {

	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi es:", pi)
	fmt.Println("Pi2 es:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values (Valores por defecto)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es:", areaCuadrado)

}
