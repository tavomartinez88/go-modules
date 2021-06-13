package interfaces

import (
	"github.com/tavomartinez88/go-modules/categories/models"
)

type ICategories interface {
	Init(name string, idRef string, availability bool)
	Build(name string, idRef string, availability bool) models.Item
}