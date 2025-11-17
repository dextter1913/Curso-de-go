package main

import "fmt"

func main() {

	// arrays
	var food [3]string
	food[0] = "hamburgesa"
	food[1] = "pizza"
	food[2] = "sushi"
	fmt.Println(food)

	//arrays literales
	food2 := [3]string{"hamburgesa", "pizza", "sushi"}

	fmt.Println(food2)
}
