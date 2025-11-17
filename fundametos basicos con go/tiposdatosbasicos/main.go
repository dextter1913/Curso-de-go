package main

import "fmt"

func main() {
	var a bool = true
	var b string = "Soy un string"
	var c rune = 123123

	// con Printf podemos imprimir datos con diferentes verbos
	// antes del verbo se debe colocar el signo de porcentaje %
	// el verbo T imprime el tipo de dato
	// el verbo v imprime el valor dentro de la variable
	// el verbo q imprime el valor con comillas
	fmt.Printf("Tipo: %T, Valor: %v ", a, a)
	fmt.Printf("Tipo: %T, Valor: %v ", b, b)
	fmt.Printf("Tipo: %T, Valor: %v ", c, c)
}
