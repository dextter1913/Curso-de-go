package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Pagado por paypal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado en efectivo")
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("Pagado por tarjeta de credito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite un numero de metodo de pago")
	fmt.Println("\t 1:paypal 2:efectivo 3:tarjeta de credito")
	_, err := fmt.Scan(&method)
	if err != nil {
		panic("debe digitar un numero valido")
	}
	if method > 3 {
		panic("debe digitar un numero valido")
	}
	PayMethod := Factory(method)
	PayMethod.Pay()
}
