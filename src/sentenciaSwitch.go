package main

import "fmt"

func main() {

	// Saber cuando un número es par o impar
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No hay condicion")
	}
}
