package main

import (
	"courses/courses"
	"fmt"
)

func main() {
	Go := courses.Course{
		Name:    "Cristian Aguirre",
		Price:   76.450,
		IsFree:  true,
		UserIDs: []uint{23, 50},
		Classes: map[uint]string{
			1: "informatica",
			2: "Español",
			3: "Ingles",
			4: "Sociales",
			5: "Desarrollo de Software",
			6: "Matematicas discretas",
		},
	}
	Go.PrintClasses(Go)
	//cambiando el valor del atributo precio usando apuntador de memoria
	Go.ChangePrice(55.90)
	fmt.Printf("el tipo es: %+v \n", Go)
	fmt.Println(Go.Price)



	// ejeplo instanciando apuntador 
	Go2 := &courses.Course{
		Name:    "Cristian Aguirre",
		Price:   76.450,
		IsFree:  true,
		UserIDs: []uint{23, 50},
		Classes: map[uint]string{
			1: "informatica",
			2: "Español",
			3: "Ingles",
			4: "Sociales",
			5: "Desarrollo de Software",
			6: "Matematicas discretas",
		},
	}
	Go2.PrintClasses(*Go2)
	fmt.Printf("el tipo es: %+v \n", Go2)
}
