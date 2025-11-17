package invoice

import (
	"pkg/pkg/customer"
	"pkg/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New turns a new invoice
func New(country string, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.GetValue()
	}
}

// GetTotal is the getter of invoice.total
func (i *Invoice) GetTotal() float64 {
	return i.total
}
