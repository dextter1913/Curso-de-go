package main

import "fmt"

// la funcion panico detiene la ejecucion del programa y nos muestra toda la traza
// del error para poder rastrear donde se genero el panico, siempre se lee desde el ultimo al primero
// es decir de abajo hacia arriba
func main() {
	division(10, 2)
	division(40, 3)
	division(12, 0)
	division(20, 4)
}

func division(dividendo, divisor int) {
	validarDivision(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivision(divisor int) {
	if divisor == 0 {
		panic("🤦")
	}
}
