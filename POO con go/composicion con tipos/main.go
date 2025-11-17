package main

import (
	"fmt"
	"pkg/pkg/customer"
	"pkg/pkg/invoice"
	"pkg/pkg/invoiceitem"
)

func main() {

	country := "Colombia"
	city := "Medellin"
	client := customer.New(
		"Cristian",
		"carrera 33 #47-35",
		"3166857000",
	)

	items := invoiceitem.NewItems(
		invoiceitem.New(1, "Mango", 3800.35),
		invoiceitem.New(1, "Mango", 3800.35),
		invoiceitem.New(1, "Mango", 3800.35),
	)

	i := invoice.New(
		country,
		city,
		client,
		items,
	)

	fmt.Printf("%+v \n", i)
	i.SetTotal()
	fmt.Println("el total es:", i.GetTotal())
}
