package builders

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/categories/models"
	"time"
)

const (
	FormatDateTimecategory = "01-02-2006 15:04:05"
)

type category struct {
	id string
	name string
	userId string
	isAvailable bool
	created string
	updated string
}

func CreateCategoryBuild() *category {
	return &category{}
}

func (c *category) Init(name string, idRef string, availability bool) {
	c.id = uuid.New().String()
	c.name = name
	c.userId = idRef
	c.isAvailable = availability
	now := time.Now()
	c.created = now.Format(FormatDateTimecategory)
	c.updated = now.Format(FormatDateTimecategory)
}

func (c *category) Build(name string, idRef string, availability bool) models.Item {
	c.Init(name, idRef, availability)
	return models.Item{
		Id: c.id,
		Name: c.name,
		RefId: c.userId,
		IsAvailable: c.isAvailable,
		IsCategory: true,
		Created: c.created,
		Updated: c.updated,
	}
}
