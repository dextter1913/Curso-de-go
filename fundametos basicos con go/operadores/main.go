package main

import "fmt"

func main() {
	// Operadores aritmenticos: (), *, /, %, +, -
	var a = 5 + (4 * 8) + 2
	var b = 4 * 2
	var c = 4 / 2
	var d = 4 % 2
	var e = 4 + 2
	var f = 4 - 2
	fmt.Println("Operadores aritmenticos: (), *, /, %, +, -")
	fmt.Printf("tipo de dato:%T valor:%v", a, a)
	fmt.Println()
	fmt.Printf("tipo de dato:%T valor:%v", b, b)
	fmt.Println()
	fmt.Printf("tipo de dato:%T valor:%v", c, c)
	fmt.Println()
	fmt.Printf("tipo de dato:%T valor:%v", d, d)
	fmt.Println()
	fmt.Printf("tipo de dato:%T valor:%v", e, e)
	fmt.Println()
	fmt.Printf("tipo de dato:%T valor:%v", f, f)
	fmt.Println()
	fmt.Println()

	// Operadores de asignacion: =, +=, -=, *=, /=, %=
	g := 4
	g += 2
	fmt.Println("Operadores de asignacion: =, +=, -=, *=, /=, %=")
	fmt.Printf("operador += tipo de dato:%T valor:%v", g, g)
	fmt.Println()
	fmt.Println()

	// Operadores de incremento y decremento: ++, --
	h := 4
	h++
	fmt.Println("Operadores de incremento y decremento: ++, --")
	fmt.Printf("operador ++ tipo de dato:%T valor:%v", h, h)
	fmt.Println()

	// Operadores de comparacion: ==, !=, <, >, <=, >=
	i := 4
	j := 2
	fmt.Println("Operadores de comparacion: ==, !=, <, >, <=, >=")
	fmt.Printf("operador == tipo de dato:%T valor:%v", i == j, i == j)
	fmt.Println()
	fmt.Println()

	// Operadores logicos: &&, ||, !
	k := 4
	l := 2
	fmt.Println("Operadores logicos: &&, ||, !")
	fmt.Printf("operador ! tipo de dato:%T valor:%v", !(k != l), !(k != l))
	fmt.Println()
	fmt.Println()
}
