package main

import (
	"fmt"
	"strings"
)

func main() {
	hello()
	hello2("Cristian", "Aguirre")

	//llamando funcion para cambiar el valor de la variable desde la memoria
	//desreferenciando, accediendo a la posicion de memoria y cambiando el valor
	emoji := "🐈"
	change(&emoji)
	fmt.Println(emoji)

	// funcion con return haciendo una sumando
	total := suma(3, 4)
	fmt.Println(total)

	// llamando funcion que retorna 2 valores
	texto := "CristIaN"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)
}

// funcion
func hello() {
	fmt.Println("Buenas 😇")
}

// funcion con parametro
func hello2(name string, lastName string) {
	fmt.Printf("Buenas %s %s", name, lastName)
}

// funcion para cambiar un valor directamente en la memoria usando apuntador de memoria
func change(value *string) {
	*value = "😋"
}

// funcion que retorna un valor con return sumando 2 valores
func suma(num1, num2 int) int {
	return num1 + num2
}

// funcion que retorna multiples valores
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}