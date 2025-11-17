package main

import "fmt"

type Storage interface {
	Get() string
	Set(string)
}

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Get() string {
	return p.Name
}

func (p *Person) Set(name string) {
	p.Name = name
}

func Exec(s Storage, name string) {
	s.Set(name)
}

func main() {
	p := NewPerson("Cristian")
	p.Set("Orlando")
	fmt.Println(p.Get())

	Exec(p, "Cristian")
	fmt.Println(p.Get())
}
