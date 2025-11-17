package main

import "fmt"

// las funciones anonimas son aquellas que no tienen nombre
func main()  {
	x := func ()  {
		fmt.Println("Hello Cristian")
	}
	x()

	//segunda forma, funcion sin variable que se autoejecute 
	func ()  {
		fmt.Println("Hello Cristian funcion autoejecutada")
	}()

	// funcion anonima autoejecutable con parametro
	func (text string)  {
		fmt.Println("Hello", text)
	}("Desde una funcion anonima autoejecutable con parametros")
}