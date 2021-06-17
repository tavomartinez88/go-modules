package methodsConcretes

import (
	"fmt"
	"github.com/tavomartinez88/go-modules/orders/delivery/interfaces"
	"github.com/tavomartinez88/go-modules/users/models/utils"
)

type deliveryM struct {
	delivery
}

func CreateToAddressDelivery(address utils.Address) interfaces.IDelivery {
	addressString := fmt.Sprintf("%#v", address)
	return &deliveryM{
		delivery{
			value: addressString,
		},
	}
}
