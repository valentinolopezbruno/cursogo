package main

// METODOS
import "fmt"

type Course struct { //CREAMOS LA ESTRUCTURA
	CourseName string
	Price      float64
	IsFree     bool
	UsersID    []uint
	Classes    map[int]string
}

//PONEMOS EL *COURSE PARA QUE LO CAMBIA DESDE LA MEMORIA DIRECTAMENTE
func (c *Course) changePrice(price float64) { // FUNC (METODO) NOMBRE (ENTRADA) SALIDA {}
	c.Price = price                                             //CAMBIAMOS PRECIO
	fmt.Println("El precio fue actualizado a: ", price, " USD") //IMPRIMIMOS
} //LUEGO LLAMAMOS AL METODO DE ESTA MANERA "(&GO).changePrice(ENTRADA)"

func main() {
	Go := Course{ //DEFINIMOS LA ESTRUCTURA CON TODOS SUS DATOS
		CourseName: "Curso de Go",
		Price:      80.32,
		IsFree:     false,
		UsersID:    []uint{1, 7, 11},
		Classes: map[int]string{
			1: "Introduccion",
			2: "Explicacion",
			3: "Examen Final",
		},
	}
	fmt.Println(Go)
	Go.changePrice(49.12) //NOTAR QUE SE PUSO (&GO) ESTO PARA HACER REFERENCIA A LA DIRECCION DE MEMORIA
	fmt.Println(Go)

}
