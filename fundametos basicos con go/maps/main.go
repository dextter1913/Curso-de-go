package main

import "fmt"

func main() {

	//mapas usando make
	animals := make(map[string]string)
	animals["cat"] = "🐈"
	animals["dog"] = "🐕"
	fmt.Println("mapa animals con make:", animals)

	// mapa literal
	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}
	fmt.Println("mapa animals2 literal:", fruits)

	//Eliminar elementos de un mapa usando la funcion delete
	delete(fruits, "banana")
	fmt.Println("frutas despues de eliminar banana:", fruits)

	//obtener el elemento del mapa de la llave cat de animals
	fmt.Println("recuperar un elemento del mapa animals:", animals["cat"])

	// cuando obtenemos el elemento del mapa nos devuelve el valor y un booleano de si existe o no la llave
	// animals["gorilla"] = "🦍"
	emoji, ok := animals["gorilla"]
	if !ok {
		animals["gorilla"] = "🦍"
		emoji, ok = animals["gorilla"]
	}
	fmt.Println("valor y si existe o no la key: ", emoji, ok)
}
