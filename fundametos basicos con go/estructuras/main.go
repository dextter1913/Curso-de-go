package main

import "fmt"

func main() {

	//creacion de la estructura
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	//instancia de la estructura con valores
	db := course{
		Name:      "Bases de datos",
		Professor: "Cristian",
		Country:   "Colombia",
	}

	//el verbo +v permite imprimir el nombre de los campos de la estructura
	fmt.Printf("%+v\n", db)

	//instancia literal la estructura, sin establecer el nombre de las keys pero si se debe respetar el orden
	db2 := course{
		"Bases de datos2",
		"Orlando",
		"Colombia",
	}

	fmt.Printf("%+v\n", db2)

	//instanciar la estructura solo con una key, en este ejemplo Proffesor
	professor := course{
		Professor: "Cristian",
	}

	fmt.Printf("%+v\n", professor)

	//ṕara acceder a cada una de las keys de la estructura, se accede como si fuera un metodo de una clase con punto (.)
	fmt.Println("name:", db.Name)
	fmt.Println("professor:", db.Professor)
	fmt.Println("country:", db.Country)

	//crear punteros a estructura
	p := &db
	p.Professor = "Aguirre"
	fmt.Printf("%+v\n", p)

}
