package director

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/categories/builders"
	"testing"
)

func TestDirector_Build_Category(t *testing.T) {
	builder := builders.CreateCategoryBuild()
	director := CreateCategoryDirector(builder)
	item := director.Build("Pizzas", "123", true)
	assert.NotNil(t, item)
	assert.NotNil(t, item.Id)
	assert.True(t, len(item.Id)>0)
	assert.True(t, item.IsCategory)
	assert.True(t, item.IsAvailable)
	assert.EqualValues(t, "Pizzas", item.Name)
	assert.EqualValues(t, "123", item.RefId)
	assert.True(t, len(item.Created)>0)
	assert.True(t, len(item.Updated)>0)
}

func TestDirector_Build_SubCategory(t *testing.T) {
	builder := builders.CreateSubCategoryBuild()
	director := CreateCategoryDirector(builder)
	item := director.Build("Pizzas", "123", true)
	assert.NotNil(t, item)
	assert.NotNil(t, item.Id)
	assert.True(t, len(item.Id)>0)
	assert.False(t, item.IsCategory)
	assert.True(t, item.IsAvailable)
	assert.EqualValues(t, "Pizzas", item.Name)
	assert.EqualValues(t, "123", item.RefId)
	assert.True(t, len(item.Created)>0)
	assert.True(t, len(item.Updated)>0)
}

