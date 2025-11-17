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
	items   invoiceitem.Items
}

// New turns a new invoice
func New(country string, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}

// GetTotal is the getter of invoice.total
func (i *Invoice) GetTotal() float64 {
	return i.total
}
