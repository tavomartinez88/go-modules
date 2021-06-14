package interfaces

import "github.com/tavomartinez88/go-modules/products/models"

type IProduct interface {
	//Init set a product initial (the parameter idRef, it refer to category or subcategory)
	Init(idRef string, name string, price float64, descriptionShort string, descriptionLarge string, hasStock bool)

	//Build generate a product with all field asigned
	Build() models.Product
}
