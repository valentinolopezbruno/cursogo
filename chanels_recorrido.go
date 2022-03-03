package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) //el numero 2 se usa para darle un limite de espacio
	c <- "Mensaje 1"          //con el <- le ordenamos la entrada al chanel
	c <- "Mebsaje 2"

	fmt.Println(len(c), cap(c)) //len es para contar y cap para mostrar la capacidad limite del chanel (en este caso len=2 y cap=2)

	// Range y Close
	// Range para recorrer cada uno de los elementos del chanel
	for message := range c {
		fmt.Println(message)
	}

	// Close se utiliza para dar aviso que el canal no recibira mas entradas
	close(c)
}
