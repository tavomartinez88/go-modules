package director

import (
	"github.com/tavomartinez88/go-modules/orders/order/interfaces"
	"github.com/tavomartinez88/go-modules/orders/order/models"
)

type director struct {
	builder interfaces.IOrder
}

func CreateOrderDirector(b interfaces.IOrder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) Build() models.Order {
	return d.builder.Build()
}