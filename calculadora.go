package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)    //declaramos el scanner
	entrada.Scan()                           //aca leemos la entrada
	operacion := entrada.Text()              //almacenamos la entrada en una variable
	fmt.Println(operacion)                   // imprimimos la variable
	valores := strings.Split(operacion, "+") // definimos otra variable para separar la entrada "2+2" por el simbolo mas y nos queda {"2","2"}
	fmt.Println(valores)
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	fmt.Println(operador1 + operador2)
}
