package director

import (
	"github.com/tavomartinez88/go-modules/categories/interfaces"
	"github.com/tavomartinez88/go-modules/categories/models"
)

type director struct {
	builder interfaces.ICategories
}

func CreateCategoryDirector(b interfaces.ICategories) *director {
	return &director{
		builder: b,
	}
}

func (d *director) Build(name string, refId string, isAvailable bool) models.Item {
	return d.builder.Build(name, refId, isAvailable)
}