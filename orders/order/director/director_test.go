package director

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/orders/order/builders"
	buildersProduct "github.com/tavomartinez88/go-modules/products/builders"
	directorProduct "github.com/tavomartinez88/go-modules/products/director"
	"github.com/tavomartinez88/go-modules/products/models"
	"testing"
)

func TestDirector_Build_with_products(t *testing.T) {
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
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"
	builderOrder.Products = []models.Product{
		product, product, product,
	}

	directorOrder := CreateOrderDirector(builderOrder)
	order := directorOrder.Build()
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

	assert.True(t, len(order.Products) == 1)
	for _, product := range order.Products {
		assert.True(t, len(product.Id)>0)
		assert.EqualValues(t, "Hamburguesa doble chesse burger", product.Name)
		assert.EqualValues(t, 520, product.Price)
		assert.EqualValues(t, 3, product.Count)
	}
}

func TestDirector_Build_without_products(t *testing.T) {
	builderOrder := builders.CreateOrderBuilder()


	builderOrder.FirstName = "Gustavo"
	builderOrder.LastName = "Martinez"
	builderOrder.PhoneNumber = "+543516143737"
	builderOrder.BirthDate = "16/11/1988"

	directorOrder := CreateOrderDirector(builderOrder)
	order := directorOrder.Build()
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

	assert.True(t, len(order.Products) == 0)
}
