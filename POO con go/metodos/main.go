package main

import (
	"fmt"
	"metodos/courses"
)

func main() {
	Go := courses.Course{
		Name:    "Cristian Aguirre",
		Price:   65.78,
		IsFree:  false,
		UserIDs: []uint{14, 90},
		Classes: map[uint]string{
			1: "instroduccion",
			2: "POO con go",
			3: "Bases de datos",
		},
	}

	fmt.Printf("Tipos de datos de la estructura: %+v \n", Go)
	Go.PrintClasses(Go)
}
