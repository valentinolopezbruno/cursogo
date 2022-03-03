package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("TENGO UNA MEMORIA DE %d GB, UN DISCO DE %d GB Y LA MARCA ES %s", myPC.ram, myPC.disk, myPC.brand)
}
func main() {
	myPC := pc{ram: 16, brand: "msi", disk: 200}
	fmt.Println(myPC)
}
