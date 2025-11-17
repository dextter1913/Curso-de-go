package main

import "fmt"

// Slice no poseen datos, son apuntadores a un array
func main() {
	set := [7]string{"🦁", "🐎", "🐄", "🦋", "🐦", "✈", "🛫"}

	//Slice seleccionado con rango, el ultimo elemento siempre se va a excluir, es decir
	//el rango final siempre sera desde la posicion 0 contando la cantidad de posiciones agregadas
	//en este caso 0,1,2,3,4 es decir, llegaria hasta el pajaro
	animals := set[0:5]
	fly := set[3:7]

	fmt.Println("Array:", set)
	fmt.Println("animales:", animals)
	fmt.Println("voladores:", fly)
	fmt.Print("\n\n")

	//los slice tienen tamaño y capacidad
	food := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}
	fruits := food[1:3]
	fruits = append(fruits, "🍍", "🍈", "🍐")

	fmt.Println("array comida:", food)
	fmt.Println("frutas:", fruits)
	fmt.Println("tamaño:", len(fruits))    //len para imprimir el tamaño del array
	fmt.Println("capacidad:", cap(fruits)) // cap para imprimir la capacidad, es decir la cantidad de indices que se le puede agregar
	fmt.Print("\n\n")

	//slice literales
	fruits2 := []string{"🍓", "🍋"}
	fruits2 = append(fruits2, "🍍")

	fmt.Println("array:", fruits2)
	fmt.Println("tamaño: ", len(fruits2))
	fmt.Println("capacidad: ", cap(fruits2))
	fmt.Print("\n\n")

	//slice usando make, recibe el slice literal, segundo parametro tamaño y tercer parametro capacidad
	fruits3 := make([]string, 0, 3)
	fruits3 = append(fruits3, "🍓", "🍋", "🍍")

	fmt.Println("array:", fruits3)
	fmt.Println("tamaño: ", len(fruits3))
	fmt.Println("capacidad: ", cap(fruits3))
	fmt.Print("\n\n")

	//valor 0 de los slice
	var fruits4 []string
	fmt.Println("array:", fruits4)
	fmt.Println("tamaño:", len(fruits4))
	fmt.Println("capacidad:", cap(fruits4))
	fmt.Println("valor nil:", fruits4 == nil)
	fmt.Print("\n\n")

	//valor 0 de los slice literal
	fruits5 := []string{}
	fmt.Println("array:", fruits5)
	fmt.Println("tamaño:", len(fruits5))
	fmt.Println("capacidad:", cap(fruits5))
	fmt.Println("valor nil:", fruits5 == nil)
	fmt.Print("\n\n")

}
