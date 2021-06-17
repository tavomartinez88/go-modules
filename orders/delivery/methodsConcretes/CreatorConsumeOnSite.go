package methodsConcretes

import "github.com/tavomartinez88/go-modules/orders/delivery/interfaces"

type onSite struct {
	consumeOnSite
}

func CreateOnSiteDelivery(v string) interfaces.IDelivery {
	return &onSite{
		consumeOnSite{
			value: v,
		},
	}
}