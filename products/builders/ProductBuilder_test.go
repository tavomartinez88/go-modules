package builders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_Build(t *testing.T) {

	builder := CreateProductBuilder()
	builder.Name = "Hamburguesa doble chesse burger"
	builder.IdRef = "1234"
	builder.Price = 520
	builder.HasStock = true
	builder.DescriptionShort = "Hamburguesa doble"
	builder.DescriptionLarge = "doble medallon de carne con panceta, pan de papas, mostaza y queso muzzarella"

	product := builder.Build()

	assert.True(t, len(product.Id)>0)
	assert.True(t, product.HasStock)
	assert.EqualValues(t, 520, product.Price)
	assert.NotNil(t, product.Id)
	assert.True(t, len(product.Id)>0)
	assert.EqualValues(t, "Hamburguesa doble chesse burger", product.Name)
	assert.EqualValues(t, "1234", product.IdRef)
	assert.EqualValues(t, "Hamburguesa doble", product.DescriptionShort)
	assert.EqualValues(t, "doble medallon de carne con panceta, pan de papas, mostaza y queso muzzarella", product.DescriptionLarge)
	assert.NotNil(t, product.Created)
	assert.True(t, len(product.Created)>0)
	assert.NotNil(t, product.Updated)
	assert.True(t, len(product.Updated)>0)
}

