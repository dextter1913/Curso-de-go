package main

import "fmt"

// apuntadores de memoria
// con *tipo dato, se le indica a go que la variable es una direccion de memoria
// para extraer el apuntador de memoria de una variable se hace con el ampersan &
// ya con eso se puede guardar en la variable de tipo apuntador de memoria en este caso la p
// si le ponemos el * a la variable podemos acceder al dato que esta en el apuntador de memoria es decir desreferenciar
func main() {
	fruit := "manzana"
	var p *string 
	p = &fruit
	*p = "pera"
	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, valor: %v,  Desreferenciacion: %s \n", p, p, *p)
}
