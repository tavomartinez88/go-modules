package director

import (
	"github.com/tavomartinez88/go-modules/products/interfaces"
	"github.com/tavomartinez88/go-modules/products/models"
)

type director struct {
	builder interfaces.IProduct
}

func CreateProductDirector(b interfaces.IProduct) *director {
	return &director{
		builder: b,
	}
}

func (d *director) Build() models.Product {
	return d.builder.Build()
}
