package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{
		id:      id,
		product: product,
		value:   value,
	}
}

func (i *Item) GetValue() float64 {
	return i.value
}
