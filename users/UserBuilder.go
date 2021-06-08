package users

type userInterface interface {
	InitUser(username string, password string)
	EncriptPassword() error
	VerifyPassword(plainPassword string) bool
	SetRole()
	GetRole() string
	SetDateTimeBuilding()
}

func GetBuilder(typeBuilder string) userInterface {
	if typeBuilder == "ADMIN" {
		return &adminBuilder{}
	}

	if typeBuilder == "PROVIDER"{
		return &providerBuilder{}
	}

	return &clientBuilder{}
}