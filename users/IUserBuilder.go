package users

type userInterface interface {
	InitUser(username string, password string)
	EncriptPassword() error
	VerifyPassword(plainPassword string) bool
	SetRole()
	GetRole() string
	SetDateTimeBuilding()
}