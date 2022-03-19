package main

import "fmt"

func main() {
	m := make(map[string]int) // Declaración de un diccionario

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	// Encontrar valor
	value, ok := m["Jose"] // Existe la llave
	fmt.Println(value, ok)
	value2, ok2 := m["Josep"] // No existe la llave por lo que se imprime 9
	fmt.Println(value2, ok2)
}
