package users

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"testing"
)

func TestProviderInitUser(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.InitUser("gustavo", "12345")
	user := providerBuilder.GetUser()
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.UserName)
	assert.NotNil(t, user.Password)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.True(t, providerBuilder.VerifyPassword("12345"))
}

func TestProviderSetRole(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetRole()
	user := providerBuilder.GetUser()
	assert.NotNil(t, user.Role)
	assert.EqualValues(t, "PROVIDER", user.Role)
}

func TestProviderGetRole(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetRole()
	role := providerBuilder.GetRole()
	assert.NotNil(t, role)
	assert.EqualValues(t, "PROVIDER", role)
}

func TestProviderSetStatus(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetStatus()
	user := providerBuilder.GetUser()
	assert.NotNil(t, user.Status)
	assert.EqualValues(t, "PENDING", user.Status)
}

func TestProviderSetDateTimeBuilding(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetDateTimeBuilding()
	user := providerBuilder.GetUser()
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
}

func TestProviderEncriptPassword(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.InitUser("gustavo", "12345")
	_ = providerBuilder.EncriptPassword()
	user := providerBuilder.GetUser()
	assert.NotNil(t, user.Password)
	assert.NotEqual(t, "12345", user.Password)
}

func TestProviderVerifyPasswordWhenIsValid(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.InitUser("gustavo", "12345")
	_ = providerBuilder.EncriptPassword()
	assert.True(t, providerBuilder.VerifyPassword("12345"))
}

func TestProviderVerifyPasswordWhenIsInValid(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.InitUser("gustavo", "12345")
	_ = providerBuilder.EncriptPassword()
	assert.False(t, providerBuilder.VerifyPassword("123456"))
}

func TestProviderClientSetAddresses(t *testing.T) {
	providerBuilder := &providerBuilder{}
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

	addresses := []utils.Address{address}

	providerBuilder.SetAddresses(addresses)

	for _, address := range providerBuilder.Addresses {
		assert.EqualValues(t,"1", address.Id)
		assert.EqualValues(t, "Av Hipolito Yrigoyen", address.Street)
		assert.EqualValues(t, "1234", address.Number)
		assert.EqualValues(t, "Center 1", address.Neigborhood)
		assert.EqualValues(t, "Venado Tuerto", address.City)
		assert.EqualValues(t, "Santa Fe", address.Province)
		assert.EqualValues(t, "2600", address.PostalCode)
		assert.EqualValues(t,"Argentina", address.Country)
		assert.False(t, address.IsDepartament)
	}
}

func TestProviderClientSetPhones(t *testing.T) {
	providerBuilder := &providerBuilder{}
	phone := utils.Phone{
		Id: "1",
		CountryCode: "+54",
		AreaCode: "3462",
		Number: "560056",
	}

	phones := []utils.Phone{phone}

	providerBuilder.SetPhones(phones)

	for _, phone := range providerBuilder.Phones {
		assert.EqualValues(t,"1", phone.Id)
		assert.EqualValues(t, "+54", phone.CountryCode)
		assert.EqualValues(t, "3462", phone.AreaCode)
		assert.EqualValues(t, "560056", phone.Number)
	}
}

func TestProviderBuilder_SetWebSite(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetWebSite("https://wwww.sarasa.com.ar")
	assert.NotNil(t, providerBuilder.WebSite)
	assert.EqualValues(t, "https://wwww.sarasa.com.ar", providerBuilder.WebSite)
}

func TestProviderBuilder_SetFacebookProfile(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetFacebookProfile("https://wwww.facebook.com/sarasa")
	assert.NotNil(t, providerBuilder.FacebookProfile)
	assert.EqualValues(t, "https://wwww.facebook.com/sarasa", providerBuilder.FacebookProfile)
}

func TestProviderBuilder_SetTwitterProfile(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetTwitterProfile("https://wwww.twitter.com/sarasa")
	assert.NotNil(t, providerBuilder.TwitterProfile)
	assert.EqualValues(t, "https://wwww.twitter.com/sarasa", providerBuilder.TwitterProfile)
}

func TestProviderBuilder_SetInstagramProfile(t *testing.T) {
	providerBuilder := &providerBuilder{}
	providerBuilder.SetInstagramProfile("https://wwww.instagram.com/sarasa")
	assert.NotNil(t, providerBuilder.InstagramProfile)
	assert.EqualValues(t, "https://wwww.instagram.com/sarasa", providerBuilder.InstagramProfile)
}

func TestProviderBuilder_SetAttentionDays(t *testing.T) {
	providerBuilder := &providerBuilder{}
	days := []utils.Day{
		{
			Id: "1",
			Name: "Lunes",
			Time: utils.Time{
				Time: utils.TimeFromTo{
					From: "09",
					To: "18",
				},
				Breaks: []utils.TimeFromTo{
					{
						From: "13",
						To: "14",
					},
				},
			},
		},
	}

	providerBuilder.SetAttentionDays(days)

	assert.NotNil(t, providerBuilder.AttentionDays)
	for _, day := range providerBuilder.AttentionDays {
		assert.EqualValues(t, "1", day.Id)
		assert.EqualValues(t, "Lunes", day.Name)
		assert.NotNil(t, day.Time)
		assert.NotNil(t, day.Time.Time)
		assert.NotNil(t, day.Time.Breaks)
	}
}


