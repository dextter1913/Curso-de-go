package courses

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// constructor para traer los atributos desde afuera
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) SetPrice(Price float64) {
	c.Price = Price
}

func (c *course) GetPrice() float64 {
	return c.Price
}

func (c *course) ListClasses() {
	text := "su clase es: "

	for _, v := range c.Classes {
		fmt.Println(text + v)
	}
}
