package courses

import "fmt"

type Course struct {
	Name     string
	price    float64
	IsFree   bool
	UserIDs  []uint
	Classess map[uint]string
}

func (c *Course) PrintClasses(data Course) {
	text := "su clase es: "
	for _, v := range data.Classess {
		fmt.Println(text + v)
	}
}

//accediendo al atributo privado price y modificando su valor o configurandolo
func (c *Course) SetPrice(Price float64) {
	c.price = Price
}

//accediendo al atributo privado price y retornando su valor
func (c *Course) GetPrice() float64 {
	price := c.price
	return price
}
