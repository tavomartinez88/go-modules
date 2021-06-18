package methodsConcretes

import (
	"github.com/tavomartinez88/go-modules/orders/delivery/interfaces"
	"github.com/tavomartinez88/go-modules/users/models/utils"
)

type deliveryM struct {
	delivery
}

const PREFIX = "Enviar a "


func CreateToAddressDelivery(address utils.Address) interfaces.IDelivery {
	var isDepartment string

	if address.IsDepartament {
		isDepartment = "Es departamento"
	} else {
		isDepartment = "No es departamento"
	}

	addressString := PREFIX+address.Street+" "+address.Number+", "+address.Neigborhood+", "+address.City+", "+address.Province+", "+address.Country+" ("+isDepartment+")"


	return &deliveryM{
		delivery{
			value: addressString,
		},
	}
}
