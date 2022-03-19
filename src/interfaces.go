package main

import "fmt"

// Estructuras

type figuras2D interface { // Aplica la función correspondiente a la figura que llegue sin necesidad de llamar el area de la figura por separado
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

// Metodos de estructuras

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// Uso de interfaz
func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	mySquare := cuadrado{base: 2}
	myRectangle := rectangulo{base: 2, altura: 4}

	calcular(mySquare)
	calcular(myRectangle)

	//Lista de interfaces
	myInterface := []interface{}{"Hola", 12, true, 3.6}
	fmt.Println(myInterface...) // ... -> Individual | Sin "..." -> Imprime la lista

}
