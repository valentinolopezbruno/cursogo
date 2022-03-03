package main

import "fmt"

// LOS STROCKS SE CREAN FUERA DE LA FUNCION DE ESTA MANERA:

type car struct {
	marca string
	año   string
}

func main() {
	// FORMA 1 DE INTERACTUAR
	autoValen := car{marca: "fox", año: "2012"}
	fmt.Println(autoValen)
	// FORMA 2 DE INTERACTUAR
	var autoSimon car
	autoSimon.marca = "cangoo"
	fmt.Println(autoSimon)

}
