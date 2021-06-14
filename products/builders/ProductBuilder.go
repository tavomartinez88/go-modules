package builders

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/products/models"
	"time"
)

const (
	FormatDateTimeProduct = "01-02-2006 15:04:05"
)

type product struct {
	IdRef string `json:"id_category"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	DescriptionShort string `json:"description_short"`
	DescriptionLarge string `json:"description_large"`
	HasStock bool `json:"has_stock"`
}

func CreateProductBuilder() *product {
	return &product{}
}

func (p *product) Init(idRef string, name string, price float64, descriptionShort string, descriptionLarge string, hasStock bool) {
	p.IdRef = idRef
	p.Name = name
	p.Price = price
	p.DescriptionShort = descriptionShort
	p.DescriptionLarge = descriptionLarge
	p.HasStock = hasStock
}

func (p *product) Build() models.Product {
	p.Init(p.IdRef, p.Name, p.Price, p.DescriptionShort, p.DescriptionLarge, p.HasStock)
	now := time.Now()
	return models.Product{
		Id: uuid.New().String(),
		Name: p.Name,
		Price: p.Price,
		DescriptionLarge: p.DescriptionLarge,
		DescriptionShort: p.DescriptionShort,
		HasStock: p.HasStock,
		IdRef: p.IdRef,
		Created: now.Format(FormatDateTimeProduct),
		Updated: now.Format(FormatDateTimeProduct),
	}
}

