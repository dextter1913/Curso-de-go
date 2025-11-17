package main

import "fmt"

func main() {
	nums := []int{1, 4, 70, 5, 67, 90, 2}

	//acá llamamos la funcion que recibe otra funcion como parametro llamada filter
	// y creamos la funcionalidad de la funcion callback para que valide si el numero es
	//menor o igual a 10 retorne verdadero, es decir en la funcion filter pasaria el if y se agregaria el valor
	//en la nueva posicion
	result := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		}
		return false
	})

	fmt.Println("resultado", result)

	//llamando la funcion que retorna otra funcion
	// ejecutamos la primera funcion y con los segundos parentecis ejecutamos la segunda funcion
	// es decir la que retornemos
	x := hello("Cristian")("Aguirre")
	fmt.Println(x)

	//otra forma de ejcutar la segunda funcion cuando tienes una funcion que retorna otra funcion
	x2 := hello("Cristian")
	fmt.Println(x2("Aguirre"))

}

// funcion qu erecibe otra funcion, parametros un slice o lista, una funcion callback el cual recibe un entero como parametro
// retorna un booleano dicha funcion call back y por ultimo retorna el listado slice de entero
func filter(nums []int, callback func(int) bool) []int {
	result := []int{}

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

// funcion que retorna una funcion
func hello(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text
	}
}
