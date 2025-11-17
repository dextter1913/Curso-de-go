package main

import "fmt"

// nos permiten recibir un numero variable de parametros del mismo tipo es decir any... la cantidad que ud desee
func main() {

	fmt.Println(suma(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// con la variable y el tipo ...tipo dato, con eso le decimos a go que el numero de patrametros es variable o dinamico
// go convertira la variable nums en una lista o slice
func suma(nums ...int) int {

	result := 0
	for _, v := range nums {
		result = result + v
	}
	return result
}
