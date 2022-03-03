package main

import "fmt"

// FUNCION MAP

func main() {
	m := make(map[string]int)

	m["simon"] = 40
	m["vvalen"] = 20
	fmt.Println(m)

	// RECORRER MAP

	for i, v := range m {
		fmt.Println(i, v)
	}

	// ENCONTRAR UN VALOR
	value, ok := m["vvalen"]
	fmt.Println(value, ok)
	//EL "OK" NOS DICE SI LA EL VALOR QUE BUSCAMOS EXISTE O NO EN EL MAPA
}
