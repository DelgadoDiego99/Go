package main

import "fmt"

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	// Array = arreglo, len = Tamaño, cap = Capacidad maxima
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])   // Imprime el indice 0
	fmt.Println(slice[:3])  // Imprime hasta el indice 3
	fmt.Println(slice[2:4]) // Imprime desde el indice 2 hasta el indice 4
	fmt.Println(slice[4:])  // Imprime desde el indice 4

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
