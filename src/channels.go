package main

import "fmt"

func say(text string, c chan<- string) { // <-chan = Salida de datos || chan<- = Entrada de datos
	c <- text
}

func main() {
	c := make(chan string, 1) // El 1 declara cuantos datos simultaneos va a manejar el canal. Si no se declara, es dinámico

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c) // Imprime todos los datos que hay en el canl

}
