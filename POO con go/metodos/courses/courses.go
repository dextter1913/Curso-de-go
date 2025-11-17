package courses

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint // Slice
	Classes map[uint]string
}

func (c Course) PrintClasses(a Course) {
	text := "las clases son: "

	for _, v := range c.Classes {
		fmt.Println(text + v)
	}
}
