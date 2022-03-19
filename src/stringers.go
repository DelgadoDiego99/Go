package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string { // Es un método de los struct para el output
	return fmt.Sprintf("Tengo %dGB de RAM, %dGB de disco y es de marca %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "Dell", disk: 256}
	myPC2 := pc{ram: 8, brand: "HP", disk: 1024}

	fmt.Println(myPC)
	fmt.Println(myPC2)
}
