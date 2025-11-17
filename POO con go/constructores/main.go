package main

import (
	"courses/courses"
	"fmt"
)

func main() {
	Go := courses.New("Go desde cero", 0, false)
	Go.UserIDs = []uint{12, 34, 54}
	Go.Classes = map[uint]string{
		1: "Español",
		2: "Etica y valores",
		3: "Ingles",
	}
	fmt.Println("el precio es:", Go.GetPrice())
	Go.ListClasses()
	Go.SetPrice(54.500)
	fmt.Printf("%+v \n", Go)
}
