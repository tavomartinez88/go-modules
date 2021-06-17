package factories

import (
	"github.com/go-errors/errors"
	"github.com/tavomartinez88/go-modules/orders/payment/interfaces"
	"github.com/tavomartinez88/go-modules/orders/payment/paymentConcretes"
)

const (
	CARD = "CARD"
	CASH = "CASH"
)

func GetPayment(typePayment string, value float64) (interfaces.IPayment, error) {

	if typePayment == ""  {
		return nil, errors.New("InfoPayment.Type is required")
	}

	if  typePayment == CARD{
		return paymentConcretes.CreateCardPayment(value), nil
	}

	if  typePayment == CASH{
		return paymentConcretes.CreateCashCreator(value), nil
	}

	return nil, errors.New("Wrong payment type passed.")
}
