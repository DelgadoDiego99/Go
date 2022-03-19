package main

import "fmt"

func main() {

	// Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "dia"
	dias := 365
	fmt.Printf("El %s dura %d dias\n", nombre, dias)
	fmt.Printf("El %v dura %v dias\n", nombre, dias)

	// Sprintf
	message := fmt.Sprintf("El %v dura %v dias", nombre, dias)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("dias: %T\n", dias)
}
