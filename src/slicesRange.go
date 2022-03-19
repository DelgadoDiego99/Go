package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {

	slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo(("ama"))         // True
	isPalindromo(("amor a roma")) // True
	isPalindromo(("casas"))       // False
	isPalindromo(("Ama"))         // True
	isPalindromo(("amA"))         // True
}
