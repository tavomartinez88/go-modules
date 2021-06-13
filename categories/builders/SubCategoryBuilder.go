package builders

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/categories/models"
	"time"
)

const (
	FormatDateTimeSubCategory = "01-02-2006 15:04:05"
)

type subCategory struct {
	id string
	name string
	categoryId string
	isAvailable bool
	created string
	updated string
}

func CreateSubCategoryBuild() *subCategory {
	return &subCategory{}
}

func (s *subCategory) Init(name string, categoryId string, availability bool) {
	s.id = uuid.New().String()
	s.name = name
	s.categoryId = categoryId
	s.isAvailable = availability
	now := time.Now()
	s.created = now.Format(FormatDateTimeSubCategory)
	s.updated = now.Format(FormatDateTimeSubCategory)
}

func (s *subCategory) Build(name string, categoryId string, availability bool) models.Item {
	s.Init(name, categoryId, availability)
	return models.Item{
		Id: s.id,
		Name: s.name,
		RefId: s.categoryId,
		IsAvailable: s.isAvailable,
		IsCategory: false,
		Created: s.created,
		Updated: s.updated,
	}
}
