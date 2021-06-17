package methodsConcretes

import "github.com/tavomartinez88/go-modules/orders/delivery/interfaces"

type AwayTake struct {
	takeAway
}

func CreateTakeAwayDelivery(v string) interfaces.IDelivery {
	return &AwayTake{
		takeAway{},
	}
}
