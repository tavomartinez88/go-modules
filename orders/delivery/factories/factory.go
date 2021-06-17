package factories

import (
	"github.com/go-errors/errors"
	"github.com/tavomartinez88/go-modules/orders/delivery/interfaces"
	"github.com/tavomartinez88/go-modules/orders/delivery/methodsConcretes"
	"github.com/tavomartinez88/go-modules/users/models/utils"
)
const TypeDeliveryOnSite = "ON_SITE"
const TypeTakeAWay = "TAKE_A_WAY"
const TypeDelivery = "DELIVERY"

func GetDelivery(typeDelivery string, description string, address ...utils.Address) (interfaces.IDelivery, error) {

	if typeDelivery == ""  {
		return nil, errors.New("InfoDelivery.Type is required")
	}

	if (typeDelivery == TypeDeliveryOnSite || typeDelivery == TypeTakeAWay) && description == ""{
		return nil, errors.New("InfoDelivery.Description is required")
	}

	if typeDelivery == TypeDelivery && address == nil {
		return nil, errors.New("InfoDelivery.Address is required")
	}

	if  typeDelivery == TypeDeliveryOnSite {
		return methodsConcretes.CreateOnSiteDelivery(description), nil
	}

	if  typeDelivery == TypeDelivery {
		return methodsConcretes.CreateToAddressDelivery(address[0]), nil
	}

	if  typeDelivery == TypeTakeAWay {
		return methodsConcretes.CreateOnSiteDelivery(description), nil
	}

	return nil, errors.New("Wrong delivery type passed.")
}