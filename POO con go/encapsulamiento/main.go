package main

import (
	"courses/courses"
	"fmt"
)

// para encapsular un metodo solo se debe poner la primera letra en minuscula
// ya con eso no se puede acceder ni al metodo ni al atributo desde afuera
func main() {
	Go := courses.Course{
		Name:    "Cristian Aguirre",
		IsFree:  false,
		UserIDs: []uint{3, 7, 9},
		Classess: map[uint]string{
			1: "Español",
			2: "Filosofia",
			3: "Sociales",
			4: "Politica",
			5: "Urbanizacion",
			6: "Desarrollo de software",
		},
	}
	Go.PrintClasses(Go)
	fmt.Printf("valores de Price: %+v \n", Go)
	Go.SetPrice(65.850)
	fmt.Printf("trayendo el valor de el Price: %v \n", Go.GetPrice())
	Go.SetPrice(20.568)
	fmt.Printf("trayendo el valor de el Price: %v \n", Go.GetPrice())
}
