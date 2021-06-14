package models

type Product struct {
	Id string `json:"id"`
	IdRef string `json:"id_ref"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	DescriptionShort string `json:"description_short"`
	DescriptionLarge string `json:"description_large"`
	HasStock bool `json:"has_stock"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
