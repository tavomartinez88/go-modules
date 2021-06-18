package factories

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"testing"
)

func TestGetDeliveryInvalidType(t *testing.T) {
	_, err := GetDelivery("", "")
	assert.NotNil(t, err)
	assert.EqualValues(t, "Type delivery is required", err.Error())
}

func TestGetDeliveryInvalidDescription(t *testing.T) {
	_, err := GetDelivery("ON_SITE", "")
	assert.NotNil(t, err)
	assert.EqualValues(t, "Description is required if type delivery is on-site or take a way", err.Error())
}

func TestGetDeliveryInvalidAddress(t *testing.T) {
	_, err := GetDelivery("DELIVERY", "")
	assert.NotNil(t, err)
	assert.EqualValues(t, "Address is required if type delivery is delivery", err.Error())
}

func TestGetDeliveryOnSiteWithDescription(t *testing.T) {
	idelivery, _ := GetDelivery("ON_SITE", "mesa 8")
	delivery := idelivery.GetDeliveryMethod()
	assert.NotNil(t, delivery)
	assert.EqualValues(t, "ON_SITE", delivery.Type)
	assert.EqualValues(t, "mesa 8", delivery.Description)
}

func TestGetDeliveryWithTypeInExistent(t *testing.T) {
	_,  err := GetDelivery("SARASA", "")
	assert.NotNil(t, err)
	assert.EqualValues(t, "Type delivery not supported", err.Error())
}

func TestGetDeliveryWithDescription(t *testing.T) {
	idelivery, _ := GetDelivery("TAKE_A_WAY", "")
	delivery := idelivery.GetDeliveryMethod()
	assert.NotNil(t, delivery)
	assert.EqualValues(t, "TAKE_A_WAY", delivery.Type)
	assert.EqualValues(t, "se retira personalmente", delivery.Description)
}

func TestGetDeliveryWithAddressNoDpto(t *testing.T) {
	address := utils.Address{
		Street: "Ernesto La Padula",
		Number: "585",
		Neigborhood: "Parque Velez Sarsfield",
		City: "Cordoba",
		Province: "Cordoba",
		PostalCode: "5000",
		Country: "ARGENTINA",
		IsDepartament: false,
	}

	idelivery, _ := GetDelivery("DELIVERY", "", address)
	delivery := idelivery.GetDeliveryMethod()
	assert.NotNil(t, delivery)
	assert.EqualValues(t, "DELIVERY", delivery.Type)
	assert.EqualValues(t, "Enviar a Ernesto La Padula 585, Parque Velez Sarsfield, Cordoba, Cordoba, ARGENTINA (No es departamento)", delivery.Description)
}

func TestGetDeliveryWithAddressDpto(t *testing.T) {
	address := utils.Address{
		Street: "Ernesto La Padula",
		Number: "585",
		Neigborhood: "Parque Velez Sarsfield",
		City: "Cordoba",
		Province: "Cordoba",
		PostalCode: "5000",
		Country: "ARGENTINA",
		IsDepartament: true,
	}

	idelivery, _ := GetDelivery("DELIVERY", "", address)
	delivery := idelivery.GetDeliveryMethod()
	assert.NotNil(t, delivery)
	assert.EqualValues(t, "DELIVERY", delivery.Type)
	assert.EqualValues(t, "Enviar a Ernesto La Padula 585, Parque Velez Sarsfield, Cordoba, Cordoba, ARGENTINA (Es departamento)", delivery.Description)
}