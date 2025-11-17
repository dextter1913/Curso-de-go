package courses

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) PrintClasses(data Course) {
	text := "su clase es :"

	for _, v := range data.Classes {
		fmt.Println(text + v)
	}
}

func (c *Course) ChangePrice(Price float64) {
	c.Price = Price
}
