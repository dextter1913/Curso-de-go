package main

import (
	"fmt"
	"os"
)

// la declaracion defer como su nombre lo indica, nos permite diferir algo
// diferir es aplazar, podemos aplazar una ejecucion hasta donde este especificado el defer
//cuando usas varios defer, debes tener en cuenta que ordenara sus ejecuciones en sentido de
// pila, es decir el primero que encuentre es el ultimo en ejecutar
// se ejecutara desde el primero que se agrego a la pila hasta el ultimo
/*
1
2
3
**/
func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	// el defer guarda el valor del momento de su ejecucion, es decir si sobre escribes la variable despues
	// de llamar el defer, cuando se ejcute la funcion al final del defer, respetara el primer valor con el
	//que se guardo el defer
	a := 5
	defer fmt.Println("Defer:", a)

	a = 10
	fmt.Println("sin el Defer:", a)

	// sus caso de uso son para cerrar conexiones, cerrar archivos, etc
	file, err := os.Create("Hello.txt")

	if err != nil {
		fmt.Printf("Ocurrio un erro al crear el archivo %v \n", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello Cristian Aguirre"))

	if err != nil {
		file.Close()
		fmt.Printf("Ocurrio un error al momento de escribir el archivo %v \n", err)
		return
	}
}
