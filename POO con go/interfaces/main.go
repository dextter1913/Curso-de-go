package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Byer()
}

type GreeterByer interface { //embebiendo interfaces
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("hola soy %s \n", p.Name)
}

func (p Person) Byer() {
	fmt.Printf("Adios soy %s \n", p.Name)
}

func (p Person) String() string {
	return "Hola soy una persona y mi nombre es " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s \n", t)
}

func (t Text) Byer() {
	fmt.Printf("Adios soy %s \n", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T \n", g, g)
	}
}

func ByerAll(bs ...Byer) {
	for _, b := range bs {
		b.Byer()
		fmt.Printf("\t Valor: %v, Tipo: %T \n", b, b)
	}
}

func All(gs ...GreeterByer) {
	for _, gb := range gs {
		gb.Greet()
		gb.Byer()
	}
}

func main() {
	var g Greeter = Person{Name: "Cristian desde Persona"}
	var x Greeter = Text("Cristian desde Text \n")

	p := Person{Name: "Critian persona 2"}
	var t Text = "Cristian en Texto 2 \n"

	l := Person{Name: "Critian persona 3"}
	var m Text = "Cristian en Texto 3 \n"
	g.Greet()
	x.Greet()
	GreetAll(p, t)

	All(l, m)

	fmt.Println(g)
}
