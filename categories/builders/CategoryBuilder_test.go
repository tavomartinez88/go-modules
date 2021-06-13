package builders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategory_Build(t *testing.T) {
	builder := CreateCategoryBuild()
	category := builder.Build("Pizzas", "123123-123", true)
	assert.NotNil(t, category)
	assert.NotNil(t, category.Id)
	assert.True(t, len(category.Id)>0)
	assert.EqualValues(t, "Pizzas", category.Name)
	assert.EqualValues(t, "123123-123", category.RefId)
	assert.True(t, len(category.Created)>0)
	assert.True(t, len(category.Updated)>0)
	assert.True(t, category.IsAvailable)
	assert.True(t, category.IsCategory)
}

