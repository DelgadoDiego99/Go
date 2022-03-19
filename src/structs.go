package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2021}
	fmt.Println(myCar)

	// Otra manera
	// Si no hay atributos iniciales, se toman los zero values
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
