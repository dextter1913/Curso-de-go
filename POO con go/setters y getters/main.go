package main

import (
	"courses/courses"
	"fmt"
)

func main() {

	Go := courses.New()
	Go.SetName("Cristian")
	Go.SetPrice(44850.23)
	Go.SetIsFree(true)
	Go.SetUserIds(1, 2, 3, 4, 5, 6)
	Go.SetClasses("español", "ingles")

	fmt.Printf("%+v \n", Go)
	fmt.Println(Go.GetUserIds())
	fmt.Println(Go.GetClasses())

	Go2 := courses.New()
	Go2.SetPrice(55400.45)

	fmt.Println("imprimiendo price Go:", Go.GetPrice())
	fmt.Println("imprimiendo price Go2:", Go2.GetPrice())

}
