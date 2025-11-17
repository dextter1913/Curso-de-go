package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//manejo de errores con go, la funcion ReadFile del paquete os, devuelve el content y el error en caso de que
	//salga un error, ese error lo podemos capturar y imprimirlo, sino hay ningun error devolveria nulo, es decir nil
	//NOTA: en go si tenemos un error controlado, continua su ejecucion, por eso debemos retornar vacio para evitar
	//que continue su ejecucion
	content, err := os.ReadFile("./manejo de errores/cualquiercosa.txt")

	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println(string(content))

	//llamando la funcion que devuelve error si trata de dividir por 0
	resultado, err := division(10, 3)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println("resultado: ", resultado)

	//llamando la funcion que devuelve error si trata de dividir por 0
	// con parametros nombrados
	resultado2, err2 := division2(10, 0)
	if err2 != nil {
		fmt.Printf("Ocurrio un error: %v", err2)
		return
	}
	fmt.Println("resultado: ", resultado2)

}

// funcion que devuelve un error en caso de que trate de dividir por 0
func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no puedes dividir por cero")
	}

	return dividendo / divisor, nil
}

// funcion que devuelve un error en caso de que trate de dividir por 0 con parametros nombrados
// con los parametros nombrados, son variables que se retornan automaticamente (NO RECOMENDABLE POR ILEGIBILIDAD)
func division2(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("no puedes dividir por cero")
		return
	}
	result = dividendo / divisor
	return
}
