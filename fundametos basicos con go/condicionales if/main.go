package main

import "fmt"

func main() {

	//condicion if
	emoji := "🌵"
	emoji = "😋"
	emoji = "🐈"

	if emoji == "🌵" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %v\n", emoji)
	}

	//declarando variable dentro del scope del if
	//se le asigna el valor en la variable dentro del if seguido de punto y coma (;)
	//emoji2 := "🐈";
	//NOTA: la variable dentro del scope del if no  sale del scope, no se pueden usar por fuera
	if emoji2 := "🐈"; emoji2 == "🌵" {
		fmt.Println("es un cactus")
	} else if emoji2 == "😋" {
		fmt.Println("Es una carita")
	} else {
		fmt.Printf("el emoji2 es: %v\n", emoji2)
	}
}
