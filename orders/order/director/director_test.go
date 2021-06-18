package director

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/orders/order/builders"
	buildersProduct "github.com/tavomartinez88/go-modules/products/builders"
	directorProduct "github.com/tavomartinez88/go-modules/products/director"
	"github.com/tavomartinez88/go-modules/products/models"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"testing"
)

func TestDirectorBuildWithoutProducts(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()


	builderOrder.FirstName = "Gustavo"
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"
	builderOrder.InfoPayment.Type = "CASH"
	builderOrder.InfoPayment.Value = 1140
	builderOrder.InfoDelivery.Type = "ON_SITE"
	builderOrder.InfoDelivery.Description = "mesa 8"

	directorOrder := CreateOrderDirector(builderOrder)
	order, _ := directorOrder.Build()
	assert.True(t, len(order.Id)>0)
	assert.EqualValues(t, "Gustavo", order.FirstName)
	assert.EqualValues(t, "Martinez", order.LastName)
	assert.EqualValues(t, "16/11/1988", order.BirthDate)
	assert.EqualValues(t, "+543516143737", order.PhoneNumber)

	assert.NotNil(t, order.Created)
	assert.True(t, len(order.Created)>0)
	assert.NotNil(t, order.Updated)
	assert.True(t, len(order.Updated)>0)
	assert.EqualValues(t, float64(0), order.Amount)
	assert.EqualValues(t, "ON_SITE", order.Delivery.GetDeliveryMethod().Type)
	assert.EqualValues(t, "mesa 8", order.Delivery.GetDeliveryMethod().Description)
	assert.True(t, len(order.Products) == 0)
}

func TestDirectorBuildWithoutTypePayment(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()

	builderOrder.FirstName = "Gustavo"
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"

	directorOrder := CreateOrderDirector(builderOrder)
	_, err := directorOrder.Build()
	assert.NotNil(t, err)
	assert.EqualValues(t, "InfoPayment.Type is required", err.Error())
}

func TestDirectorBuildWithProductsAndPaymentValid(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()
	builder := buildersProduct.CreateProductBuilder()

	builder.Name = "Hamburguesa doble chesse burger"
	builder.IdRef = "1234"
	builder.Price = 520
	builder.HasStock = true
	builder.DescriptionShort = "Hamburguesa doble"
	builder.DescriptionLarge = "doble medallon de carne con panceta, pan de papas, mostaza y queso muzzarella"
	productDirector := directorProduct.CreateProductDirector(builder)
	product := productDirector.Build()

	builderOrder.FirstName = "Gustavo"
	builderOrder.InfoPayment.Type = "CASH"
	builderOrder.InfoPayment.Value = 1140
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"
	builderOrder.InfoDelivery.Type = "DELIVERY"
	builderOrder.InfoDelivery.Address = utils.Address{
		Street: "Ernesto La Padula",
		Number: "585",
		Neigborhood: "Parque Velez Sarsfield",
		City: "Cordoba",
		Province: "Cordoba",
		PostalCode: "5000",
		Country: "ARGENTINA",
		IsDepartament: true,
	}

	builderOrder.Products = []models.Product{
		product, product, product,
	}

	directorOrder := CreateOrderDirector(builderOrder)
	order, _ := directorOrder.Build()
	assert.True(t, len(order.Id)>0)
	assert.EqualValues(t, "Gustavo", order.FirstName)
	assert.EqualValues(t, "Martinez", order.LastName)
	assert.EqualValues(t, "16/11/1988", order.BirthDate)
	assert.EqualValues(t, "+543516143737", order.PhoneNumber)

	assert.NotNil(t, order.Created)
	assert.True(t, len(order.Created)>0)
	assert.NotNil(t, order.Updated)
	assert.True(t, len(order.Updated)>0)
	assert.EqualValues(t, float64(1560), order.Amount)
	assert.EqualValues(t, "CASH", order.Payment.GetPayment().Type)
	assert.EqualValues(t, 1140, order.Payment.GetPayment().CountMoneyToPay)
	assert.EqualValues(t, 0, order.Payment.GetPayment().PercentageForCard)
	assert.EqualValues(t, "DELIVERY", order.Delivery.GetDeliveryMethod().Type)
	assert.EqualValues(t, "Enviar a Ernesto La Padula 585, Parque Velez Sarsfield, Cordoba, Cordoba, ARGENTINA (Es departamento)", order.Delivery.GetDeliveryMethod().Description)
	assert.True(t, len(order.Products) == 1)
	for _, product := range order.Products {
		assert.True(t, len(product.Id)>0)
		assert.EqualValues(t, "Hamburguesa doble chesse burger", product.Name)
		assert.EqualValues(t, 520, product.Price)
		assert.EqualValues(t, 3, product.Count)
	}
}

func TestDirectorBuildWithProductsAndPaymentInValid(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()
	builder := buildersProduct.CreateProductBuilder()

	builder.Name = "Hamburguesa doble chesse burger"
	builder.IdRef = "1234"
	builder.Price = 520
	builder.HasStock = true
	builder.DescriptionShort = "Hamburguesa doble"
	builder.DescriptionLarge = "doble medallon de carne con panceta, pan de papas, mostaza y queso muzzarella"
	productDirector := directorProduct.CreateProductDirector(builder)
	product := productDirector.Build()

	builderOrder.FirstName = "Gustavo"
	builderOrder.InfoPayment.Type = "CARD"
	builderOrder.InfoPayment.Value = 10.5
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"
	builderOrder.InfoDelivery.Type = "ON_SITE"
	builderOrder.InfoDelivery.Description = "mesa 8"
	builderOrder.Products = []models.Product{
		product, product, product,
	}

	directorOrder := CreateOrderDirector(builderOrder)
	order, _ := directorOrder.Build()
	assert.True(t, len(order.Id)>0)
	assert.EqualValues(t, "Gustavo", order.FirstName)
	assert.EqualValues(t, "Martinez", order.LastName)
	assert.EqualValues(t, "16/11/1988", order.BirthDate)
	assert.EqualValues(t, "+543516143737", order.PhoneNumber)

	assert.NotNil(t, order.Created)
	assert.True(t, len(order.Created)>0)
	assert.NotNil(t, order.Updated)
	assert.True(t, len(order.Updated)>0)
	assert.EqualValues(t, float64(1560), order.Amount)
	assert.EqualValues(t, "CARD", order.Payment.GetPayment().Type)
	assert.EqualValues(t, 0, order.Payment.GetPayment().CountMoneyToPay)
	assert.EqualValues(t, 10.5, order.Payment.GetPayment().PercentageForCard)

	assert.True(t, len(order.Products) == 1)
	for _, product := range order.Products {
		assert.True(t, len(product.Id)>0)
		assert.EqualValues(t, "Hamburguesa doble chesse burger", product.Name)
		assert.EqualValues(t, 520, product.Price)
		assert.EqualValues(t, 3, product.Count)
	}
}

func TestDirectorBuildWithoutDelivery(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()
	builder := buildersProduct.CreateProductBuilder()

	builder.Name = "Hamburguesa doble chesse burger"
	builder.IdRef = "1234"
	builder.Price = 520
	builder.HasStock = true
	builder.DescriptionShort = "Hamburguesa doble"
	builder.DescriptionLarge = "doble medallon de carne con panceta, pan de papas, mostaza y queso muzzarella"
	productDirector := directorProduct.CreateProductDirector(builder)
	product := productDirector.Build()

	builderOrder.FirstName = "Gustavo"
	builderOrder.InfoPayment.Type = "CARD"
	builderOrder.InfoPayment.Value = 10.5
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"
	builderOrder.Products = []models.Product{
		product, product, product,
	}

	directorOrder := CreateOrderDirector(builderOrder)
	_, err := directorOrder.Build()
	assert.NotNil(t, err)
	assert.EqualValues(t, "Type delivery is required", err.Error())
}
