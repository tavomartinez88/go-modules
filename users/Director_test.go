package users

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"testing"
)

func TestAdmin_buildUser(t *testing.T) {
	adminBuilder := newAdminBuilder()
	adminBuilder.SetUsername("gustavo")
	adminBuilder.SetPassword("martinez")
	director := newDirector(adminBuilder)
	user := director.buildUser()
	assert.NotNil(t, user)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.NotNil(t, user.Password)
	assert.True(t, director.builder.VerifyPassword("martinez"))
	assert.EqualValues(t, "ADMIN", user.Role)
	assert.Nil(t, user.Addresses)
	assert.EqualValues(t, "", user.WebSite)
	assert.EqualValues(t, "", user.FacebookProfile)
	assert.EqualValues(t, "", user.TwitterProfile)
	assert.EqualValues(t, "", user.InstagramProfile)
	assert.Nil(t, user.Phones)
	assert.EqualValues(t, "ACTIVE", user.Status)
	assert.Nil(t, user.AttentionDays)
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
	assert.True(t, len(user.Created)>0)
	assert.NotNil(t, len(user.Updated)>0)
}

func TestClient_buildUser(t *testing.T) {
	clientBuilder := newClientBuilder()
	clientBuilder.SetUsername("gustavo")
	clientBuilder.SetPassword("martinez")
	clientBuilder.Address = utils.Address{
		Id: "1",
		Street: "Av siempre viva",
		Number: "123",
		Neigborhood: "barrio 1",
		City: "Sprinfield",
		Province: "Spr",
		PostalCode: "1234",
		Country: "USA",
		IsDepartament: false,
	}

	clientBuilder.Phone = utils.Phone{
		Id: "1",
		CountryCode: "+54",
		Number: "560056",
		AreaCode: "3462",
	}

	director := newDirector(clientBuilder)
	user := director.buildUser()

	assert.NotNil(t, clientBuilder)
	assert.NotNil(t, user)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.NotNil(t, user.Password)
	assert.True(t, director.builder.VerifyPassword("martinez"))
	assert.EqualValues(t, "CLIENT", user.Role)

	assert.NotNil(t, user.Addresses)
	if user.Addresses != nil {
		for _, address := range user.Addresses {
			assert.EqualValues(t, "1", address.Id)
			assert.EqualValues(t, "Av siempre viva", address.Street)
			assert.EqualValues(t, "barrio 1", address.Neigborhood)
			assert.EqualValues(t, "123", address.Number)
			assert.EqualValues(t, "1234", address.PostalCode)
			assert.EqualValues(t, "Spr", address.Province)
			assert.EqualValues(t, "Sprinfield", address.City)
			assert.EqualValues(t, "USA", address.Country)
			assert.False(t, address.IsDepartament)
		}
	}

	if user.Phones != nil {
		for _, phone := range user.Phones {
			assert.EqualValues(t, "1", phone.Id)
			assert.EqualValues(t, "+54", phone.CountryCode)
			assert.EqualValues(t, "3462", phone.AreaCode)
			assert.EqualValues(t, "560056", phone.Number)
		}
	}

	assert.EqualValues(t, "", user.WebSite)
	assert.EqualValues(t, "", user.FacebookProfile)
	assert.EqualValues(t, "", user.TwitterProfile)
	assert.EqualValues(t, "", user.InstagramProfile)

	assert.NotNil(t, user.Phones)

	assert.Nil(t, user.AttentionDays)
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
	assert.True(t, len(user.Created)>0)
	assert.NotNil(t, len(user.Updated)>0)
}


func TestProvider_buildUser(t *testing.T) {
	providerBuilder := newProviderBuilder()
	providerBuilder.SetUsername("gustavo")
	providerBuilder.SetPassword("martinez")
	providerBuilder.SetAddresses([]utils.Address{
		{
			Id: "1",
			Street: "Av siempre viva",
			Number: "123",
			Neigborhood: "barrio 1",
			City: "Sprinfield",
			Province: "Spr",
			PostalCode: "1234",
			Country: "USA",
			IsDepartament: false,
		},
	})

	providerBuilder.SetPhones([]utils.Phone{
		{
			Id: "1",
			CountryCode: "+54",
			Number: "560056",
			AreaCode: "3462",
		},
	})

	providerBuilder.SetWebSite("https://www.miweb.com")
	providerBuilder.SetFacebookProfile("https://www.miweb.com")
	providerBuilder.SetInstagramProfile("https://www.miweb.com")
	providerBuilder.SetTwitterProfile("https://www.miweb.com")

	director := newDirector(providerBuilder)
	user := director.buildUser()

	assert.NotNil(t, providerBuilder)
	assert.NotNil(t, user)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.NotNil(t, user.Password)
	assert.True(t, director.builder.VerifyPassword("martinez"))
	assert.EqualValues(t, "PROVIDER", user.Role)

	assert.NotNil(t, user.Addresses)
	if user.Addresses != nil {
		for _, address := range user.Addresses {
			assert.EqualValues(t, "1", address.Id)
			assert.EqualValues(t, "Av siempre viva", address.Street)
			assert.EqualValues(t, "barrio 1", address.Neigborhood)
			assert.EqualValues(t, "123", address.Number)
			assert.EqualValues(t, "1234", address.PostalCode)
			assert.EqualValues(t, "Spr", address.Province)
			assert.EqualValues(t, "Sprinfield", address.City)
			assert.EqualValues(t, "USA", address.Country)
			assert.False(t, address.IsDepartament)
		}
	}

	if user.Phones != nil {
		for _, phone := range user.Phones {
			assert.EqualValues(t, "1", phone.Id)
			assert.EqualValues(t, "+54", phone.CountryCode)
			assert.EqualValues(t, "3462", phone.AreaCode)
			assert.EqualValues(t, "560056", phone.Number)
		}
	}

	assert.EqualValues(t, "https://www.miweb.com", user.WebSite)
	assert.EqualValues(t, "https://www.miweb.com", user.FacebookProfile)
	assert.EqualValues(t, "https://www.miweb.com", user.TwitterProfile)
	assert.EqualValues(t, "https://www.miweb.com", user.InstagramProfile)

	assert.NotNil(t, user.Phones)

	assert.Nil(t, user.AttentionDays)
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
	assert.True(t, len(user.Created)>0)
	assert.NotNil(t, len(user.Updated)>0)
}