package main

import (
	"fmt"
)

type Course struct {
	Name    string
	Price   float64
	isFree  bool
	UserIds []uint //Slice
	classes map[uint]string
}

func main() {

	// instanciando la estructura a una variable
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		isFree:  false,
		UserIds: []uint{12, 56, 89},
		classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	fmt.Printf("estructura es de tipo: %T con un valor %+v \n", Go)

	//podemos ingresar los datos a la estructura sin necesidad de declarar los nombres de los atributos
	// el requisito es hacerlo en el mismo orden el cual estan creados en la estructura
	Go2 := Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}
	fmt.Printf("estructura es de tipo: %T con un valor %v \n", Go2.Name, Go2.Name)
	fmt.Printf("estructura es de tipo: %+v \n", Go2)

	css := Course{
		Name:   "css desde 0",
		isFree: true,
	}

	fmt.Println("nombre del curso:", css.Name)
	fmt.Printf("lista de campos de la variable css %+v \n", css)

	//otra forma de asignar valores sin llamar los campos
	js := Course{}
	js.Name = "Curso de javascript"
	js.UserIds = []uint{12, 67}

	fmt.Println("nombre del curso:", js.Name)
	fmt.Printf("lista de campos de la variable javascript %+v \n", js)

}
