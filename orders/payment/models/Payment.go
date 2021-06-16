package models

//CountMoneyToPay amount money to pay
//PercentageForCard represent percentage rechage if the payment's type is card else if amount chash
type Payment struct {
	Type string
	CountMoneyToPay float64
	PercentageForCard float64
}
