package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

// acá podemos embeber la estructura persona dentro de la estructura employee y ya llamando employee accedemos a los metodos de persona
type Employee struct {
	Person
	Salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {

	e := NewEmployee("Cristian", 30, 1000000)

	e.Payroll()

	fmt.Println(e.Name)
	fmt.Println(e.Age)
	e.Greet()
}
