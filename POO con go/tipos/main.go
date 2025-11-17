package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declaracion de alias
type myAlias = course

// definicion de tipo (no hereda sus metodos, solo los campos de la estructura)
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "Verdadero"
	}
	return "falso"
}

func main() {
	a := myAlias{
		name: "Go",
	}
	a.Print()
	fmt.Printf("el tipo de a es: %T \n", a)

	b := newCourse{
		name: "Go2",
	}

	fmt.Printf("el tipo de b es: %T \n", b)

	var c newBool = false
	fmt.Printf("%+v \n", c.String())
}
