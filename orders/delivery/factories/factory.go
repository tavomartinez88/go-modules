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

func GetDelivery(typeDelivery string, description string, addresses ...utils.Address) (interfaces.IDelivery, error) {

	if typeDelivery == ""  {
		return nil, errors.New("Type delivery is required")
	}

	if (typeDelivery == TypeDeliveryOnSite) && description == "" {
		return nil, errors.New("Description is required if type delivery is on-site or take a way")
	}

	if typeDelivery == TypeDelivery && addresses == nil {
		return nil, errors.New("Address is required if type delivery is delivery")
	}

	if  typeDelivery == TypeDeliveryOnSite {
		return methodsConcretes.CreateOnSiteDelivery(description), nil
	}

	if  typeDelivery == TypeDelivery {
		return methodsConcretes.CreateToAddressDelivery(addresses[0]), nil
	}

	if  typeDelivery == TypeTakeAWay {
		return methodsConcretes.CreateTakeAwayDelivery(), nil
	}

	return nil, errors.New("Type delivery not supported")
}