package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashPayment(t *testing.T) {
	payment, _ := GetPayment("CASH", 565)
	assert.NotNil(t, payment)
	assert.EqualValues(t, "CASH", payment.GetPayment().Type)
	assert.EqualValues(t, 565, payment.GetPayment().CountMoneyToPay)
}

func TestGetCardPayment(t *testing.T) {
	payment, _ := GetPayment("CARD", 15.5)
	assert.NotNil(t, payment)
	assert.EqualValues(t, "CARD", payment.GetPayment().Type)
	assert.EqualValues(t, 15.5, payment.GetPayment().PercentageForCard)
}

func TestGetCryptoPayment(t *testing.T) {
	_, err := GetPayment("CRYPTO", 150)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Wrong payment type passed.", err.Error())
}
