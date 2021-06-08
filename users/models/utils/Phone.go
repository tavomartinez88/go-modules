package utils

//todo agregar validacion para verificar que es un numero telefonico
type Phone struct {
	Id string `json:"id"`
	CountryCode string `json:"country_code" validate:"required`
	AreaCode string `json:"area_code" validate:"required`
	Number string `json:"number" validate:"required`
}
