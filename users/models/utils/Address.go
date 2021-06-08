package utils

type Address struct {
	Id string `json:"id"`
	Street string `json:"street" validate:"required, min=3, max=255"`
	Number string `json:"number" validate:"required, min=3, max=255"`
	Neigborhood string `json:"number" validate:"required, min=3, max=50"`
	City string `json:"number" validate:"required, min=2, max=50"`
	Province string `json:"number" validate:"required, min=3, max=50"`
	PostalCode string `json:"number" validate:"required, min=2, max=10"`
	Country string `json:"number" validate:"required, min=3, max=255"`
	IsDepartament bool `json:"is_departament"`
}
