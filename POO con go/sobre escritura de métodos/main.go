package main

import "fmt"

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age uint, children uint) Human {
	return Human{age, children}
}

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hola desde persona")
}

// acá podemos embeber la estructura persona dentro de la estructura employee y ya llamando employee accedemos a los metodos de persona
type Employee struct {
	Person
	Human
	Salary float64
}

func NewEmployee(name string, age uint, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

// sobre escritura del metodo saludar de Persona
func (e Employee) Greet() {
	fmt.Println("Hola desde empleado")
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {

	e := NewEmployee("Cristian", 30, 2, 1000000)

	e.Payroll()
	fmt.Println(e.Name)
	fmt.Println(e.Children)
	fmt.Println(e.Human.Age)
	e.Greet()
	e.Person.Greet()
}
