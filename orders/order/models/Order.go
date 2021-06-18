package models

import (
	interfaceDelivery "github.com/tavomartinez88/go-modules/orders/delivery/interfaces"
	"github.com/tavomartinez88/go-modules/orders/payment/interfaces"
)

type Order struct {
	Id string `json:"id"`
	Amount float64 `json:"amount"`
	Products []ProductItem
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Payment interfaces.IPayment
	Delivery interfaceDelivery.IDelivery
	Created string `json:"created"`
	Updated string `json:"updated"`
}

//ProductItem is an struct with info more important
type ProductItem struct {
	Id string `json:"id"`
	Price float64 `json:"price"`
	Count int `json:"count"`
	Name string `json:"name"`
}


