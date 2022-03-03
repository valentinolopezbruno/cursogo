package main

import "fmt"

func par(a int) int {
	return a % 2
}

func main() {

	value := par(3)

	if value == 0 {

		fmt.Println("el numero es par")
	} else {
		fmt.Println("el numero es impar")
	}
}
