package main


import "fmt"

// la funcion recover sirve para recuperarnos despues de un panic
func main() {
	division(10, 2)
	division(40, 3)
	division(12, 0)
	division(20, 4)
}

func division(dividendo, divisor int) {
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Recuperandome del panic", r)
		}

	}()
	validarDivision(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivision(divisor int) {
	if divisor == 0 {
		panic("🤦")
	}
}
