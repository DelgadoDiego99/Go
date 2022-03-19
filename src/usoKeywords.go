package main

import "fmt"

func main() {
	// Defer (Ejecuta la linea de ultimo) Ej: Cerrar sesion/conexion/archivo
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue (Continua el codigo apesar de un error (Controlado))
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
