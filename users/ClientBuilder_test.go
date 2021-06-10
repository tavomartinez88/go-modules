package users

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"testing"
)

func TestClientInitUser(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.InitUser("gustavo", "12345")
	user := clientBuilder.GetUser()
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.UserName)
	assert.NotNil(t, user.Password)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.EqualValues(t, "12345", user.Password)
}

func TestClientSetRole(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.SetRole()
	user := clientBuilder.GetUser()
	assert.NotNil(t, user.Role)
	assert.EqualValues(t, "CLIENT", user.Role)
}

func TestClientGetRole(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.SetRole()
	role := clientBuilder.GetRole()
	assert.NotNil(t, role)
	assert.EqualValues(t, "CLIENT", role)
}

func TestClientSetDateTimeBuilding(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.SetDateTimeBuilding()
	user := clientBuilder.GetUser()
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
}

func TestClientEncriptPassword(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.InitUser("gustavo", "12345")
	_ = clientBuilder.EncriptPassword()
	user := clientBuilder.GetUser()
	assert.NotNil(t, user.Password)
	assert.NotEqual(t, "12345", user.Password)
}

func TestClientVerifyPasswordWhenIsValid(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.InitUser("gustavo", "12345")
	clientBuilder.EncriptPassword()
	assert.True(t, clientBuilder.VerifyPassword("12345"))
}

func TestClientVerifyPasswordWhenIsInValid(t *testing.T) {
	clientBuilder := &clientBuilder{}
	clientBuilder.InitUser("gustavo", "12345")
	clientBuilder.EncriptPassword()
	assert.False(t, clientBuilder.VerifyPassword("123456"))
}

func TestClientSetAddress(t *testing.T) {
	address := utils.Address{
		Id: "1",
		Street: "Av Hipolito Yrigoyen",
		Number: "1234",
		Neigborhood: "Center 1",
		City: "Venado Tuerto",
		Province: "Santa Fe",
		PostalCode: "2600",
		Country: "Argentina",
		IsDepartament: false,
	}

	clientBuilder := &clientBuilder{}
	clientBuilder.SetAddress(address)
	assert.NotNil(t, clientBuilder.Address)
	assert.EqualValues(t,"1", clientBuilder.Address.Id)
	assert.EqualValues(t, "Av Hipolito Yrigoyen", clientBuilder.Address.Street)
	assert.EqualValues(t, "1234", clientBuilder.Address.Number)
	assert.EqualValues(t, "Center 1", clientBuilder.Address.Neigborhood)
	assert.EqualValues(t, "Venado Tuerto", clientBuilder.Address.City)
	assert.EqualValues(t, "Santa Fe", clientBuilder.Address.Province)
	assert.EqualValues(t, "2600", clientBuilder.Address.PostalCode)
	assert.EqualValues(t,"Argentina", clientBuilder.Address.Country)
	assert.False(t, clientBuilder.Address.IsDepartament)
}

func TestClientBuilder_SetPhone(t *testing.T) {
	phone := utils.Phone{
		Id: "1",
		CountryCode: "+54",
		AreaCode: "3462",
		Number: "560056",
	}

	clientBuilder := &clientBuilder{}
	clientBuilder.SetPhone(phone)
	assert.NotNil(t, clientBuilder.Phone)
	assert.EqualValues(t,"1", clientBuilder.Phone.Id)
	assert.EqualValues(t, "+54", clientBuilder.Phone.CountryCode)
	assert.EqualValues(t, "3462", clientBuilder.Phone.AreaCode)
	assert.EqualValues(t, "560056", clientBuilder.Phone.Number)
}




