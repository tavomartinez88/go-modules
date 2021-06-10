package users

import "github.com/tavomartinez88/go-modules/users/models"

type director struct {
	builder userInterface
}

func newDirector(b userInterface) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildUser() models.User {
	return d.builder.GetUser()
}