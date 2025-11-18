package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {

	v, ok := i.(string) //implementando type assertions, es decir validando el tipo que se guarda en la interface vacia
	if ok {
		fmt.Println(strings.ToUpper(v))
	}

	switch v := i.(type) { // implementando type switches, es decir valida el tipo dentro de un switch
	case string:
		fmt.Println(strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("valor: %v, tipo: %T \n", i, i)
	}
}

func main() {
	wrapper(12)
	wrapper(14.3)
	wrapper(false)
	wrapper("Hola mundo")
	// var e exampler

	// fmt.Printf("valor: %v, tipo: %T \n", e, e)

	// e.x()
}
