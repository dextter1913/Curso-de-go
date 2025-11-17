package main

import "fmt"

func main() {
	// ciclo for clasico, la misma sintaxis de otros lenguajes con la diferencia de que no es necesario los parentecis ()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// ciclo for continuo similar a una sintaxis de ciclo while en otros lenguajes de programacion
	i2 := 1
	for i2 <= 10 {
		fmt.Println(i2)
		i2++
	}

	fmt.Println()

	//ciclo for forever es un for que dura para siempre
	// i3 := 1
	// for {
	// 	fmt.Println(i3)
	// 	i3++
	// }

	//ciclo for range slice, sirve para iterar sobre slices, mapas y strings
	//imprimiendo sobre slice usando logica clasica
	nums := []uint8{2, 4, 6, 8}
	i4 := 0

	for range nums {
		fmt.Println("Imprimiendo cositas", nums[i4])
		i4++
	}

	//imprimiendo sobre slice usando logica de range con indice y valor (i5, v5)
	for i5, v5 := range nums {
		fmt.Println("indice: ", i5, "Valor: ", v5)
	}

	//imprimiendo valores del slice usando ciclo for range multiplicando valores por *2
	for i5 := range nums {
		nums[i5] *= 2
	}
	fmt.Println(string(nums))

	// iterar sobre mapas usando for range
	sports := map[string]string{"basketball": "🏀", "soccer": "⚽", "baseball": "⚾"}
	fmt.Printf("tipo: %T \nvalores: %v \n", sports, sports)
	for key, v6 := range sports {
		fmt.Println("Key:", key, "value:", v6)
	}

	//iterar sobre string usando for range, recuerda que si no lo casteas como string
	//se imprimen los bytes mas no el string que es lo que necesitamos en este caso con
	//la funcion string()
	hello := "hello"

	for _, v7 := range hello {
		fmt.Println(string(v7))
	}
}
