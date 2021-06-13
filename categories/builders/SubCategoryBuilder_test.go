package builders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubCategory_Build(t *testing.T) {
	builder := CreateSubCategoryBuild()
	category := builder.Build("Pizzas con salsa roja", "123123-123", true)
	assert.NotNil(t, category)
	assert.NotNil(t, category.Id)
	assert.True(t, len(category.Id)>0)
	assert.EqualValues(t, "Pizzas con salsa roja", category.Name)
	assert.EqualValues(t, "123123-123", category.RefId)
	assert.True(t, len(category.Created)>0)
	assert.True(t, len(category.Updated)>0)
	assert.True(t, category.IsAvailable)
	assert.False(t, category.IsCategory)
}

