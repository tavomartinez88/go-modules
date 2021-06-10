package users

import "github.com/tavomartinez88/go-modules/users/models"

type userInterface interface {
	InitUser(username string, password string)
	EncriptPassword() error
	VerifyPassword(plainPassword string) bool
	SetUsername(username string)
	SetPassword(password string)
	SetRole()
	GetUser() models.User
	GetRole() string
	SetDateTimeBuilding()
}