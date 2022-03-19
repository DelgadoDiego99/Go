package main

import (
	"fmt"
	"log"     // Log for errors
	"strconv" // Parser
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("El valor es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("True AND")
	} else {
		fmt.Println("False AND")
	}

	// OR
	if valor1 == 0 && valor2 == 2 {
		fmt.Println("True OR")
	} else {
		fmt.Println("False OR")
	}

	// Convertir texto a número
	value, err := strconv.Atoi("53") // Atoi = ParseInt but to String
	if err != nil {                  // Se usa para conocer si hay errores
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
