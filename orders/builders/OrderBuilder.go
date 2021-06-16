package builders

import (
	"github.com/google/uuid"
	model "github.com/tavomartinez88/go-modules/orders/models"
	"github.com/tavomartinez88/go-modules/products/models"
	"time"
)

const (
	FormatDateTimeProduct = "01-02-2006 15:04:05"
)

type order struct {
	id string
	amount float64
	Products []models.Product
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	created string
	updated string
}

func CreateOrderBuilder() *order {
	return &order{}
}

func (o *order) buildProductItem(product models.Product) model.ProductItem {
	return model.ProductItem{
		Id: product.Id,
		Price: product.Price,
		Name: product.Name,
		Count: o.getCountProducts(product),
	}
}

func (o *order) getCountProducts(product models.Product) int {
	var timesFound = 0
	for _, p := range o.Products {
		if p.Id == product.Id && p.Name == product.Name {
			timesFound++
		}
	}

	return timesFound
}

func contains(items []model.ProductItem, item model.ProductItem) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}

func (o *order) getListBuildProductItem() []model.ProductItem{
	var productItems = []model.ProductItem{}
	for _, product := range o.Products {
		productItem := o.buildProductItem(product)
		if !contains(productItems, productItem) {
			productItems = append(productItems, productItem)
		}
	}

	return productItems
}

func (o *order) Init() {
	var totalPrice float64 = 0
	o.id = uuid.New().String()

	for _, product := range o.Products  {
		totalPrice = totalPrice + product.Price
	}

	o.amount = totalPrice

	now := time.Now()
	o.created = now.Format(FormatDateTimeProduct)
	o.updated = now.Format(FormatDateTimeProduct)
}

func (o *order) Build() model.Order {
	o.Init()
	return model.Order{
		Id: o.id,
		Amount: o.amount,
		Products: o.getListBuildProductItem(),
		FirstName: o.FirstName,
		LastName: o.LastName,
		BirthDate: o.BirthDate,
		PhoneNumber: o.PhoneNumber,
		Created: o.created,
		Updated: o.updated,
	}
}

