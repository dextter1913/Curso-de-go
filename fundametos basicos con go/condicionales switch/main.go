package main

import "fmt"

func main() {

	//condicional switch, no es necesario parentecis y break como en otros lenguajes
	emoji := "🐈"
	emoji = "🍌"
	emoji = "🐎"
	switch emoji {
	case "🐈":
		fmt.Println("el emoji es un gato")
	case "🐕":
		fmt.Println("el emoji es un perro")
	case "🍌":
		fmt.Println("el emoji es una banana")
	case "🍎":
		fmt.Println("el emoji es una manzana")
	default:
		fmt.Printf("el emoji no es ni animal ni fruta: %s\n", emoji)
	}

	//con go podemos agrupar casos con comas (,)
	emoji2 := "🐦"
	emoji2 = "🍓"
	emoji2 = "🛫"
	switch emoji2 {
	case "🐎", "🐈", "🐦":
		fmt.Println("es un animal")
	case "🍋", "🍍", "🍓", "🍎":
		fmt.Println("es una fruta")
	default:
		fmt.Printf("el emoji2 no es ni animal ni fruta: %s\n", emoji2)
	}

	//con go tambien podemos agregar operadores de comparacion y operadores logicos,
	// solo debemos omitir la expresion del switch es decir switch { case variable operador logico o comparacion}
	// ejemplo
	emoji3 := "🐺"
	emoji3 = "🛬"
	switch {
	case emoji3 == "🐺" || emoji3 == "🐦":
		fmt.Println("es un animal")
	case emoji3 == "🍐" || emoji == "🍈":
		fmt.Println("es una fruta")
	default:
		fmt.Printf("el emoji3 no es ni animal ni fruta: %s", emoji3)
	}
}
