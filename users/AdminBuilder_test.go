package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitUser(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.InitUser("gustavo", "12345")
	user := adminBuilder.GetUser()
	assert.NotNil(t, user.Id)
	assert.NotNil(t, user.UserName)
	assert.NotNil(t, user.Password)
	assert.EqualValues(t, "gustavo", user.UserName)
	assert.EqualValues(t, "12345", user.Password)
}

func TestSetRole(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.SetRole()
	user := adminBuilder.GetUser()
	assert.NotNil(t, user.Role)
	assert.EqualValues(t, "ADMIN", user.Role)
}

func TestGetRole(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.SetRole()
	role := adminBuilder.GetRole()
	assert.NotNil(t, role)
	assert.EqualValues(t, "ADMIN", role)
}

func TestSetStatus(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.SetStatus()
	user := adminBuilder.GetUser()
	assert.NotNil(t, user.Status)
	assert.EqualValues(t, "ACTIVE", user.Status)
}

func TestSetDateTimeBuilding(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.SetDateTimeBuilding()
	user := adminBuilder.GetUser()
	assert.NotNil(t, user.Created)
	assert.NotNil(t, user.Updated)
}

func TestEncriptPassword(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.EncriptPassword()
	user := adminBuilder.GetUser()
	assert.NotNil(t, user.Password)
	assert.NotEqual(t, "12345", user.Password)
}

func TestVerifyPasswordWhenIsValid(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.InitUser("gustavo", "12345")
	_ = adminBuilder.EncriptPassword()
	assert.True(t, adminBuilder.VerifyPassword("12345"))
}

func TestVerifyPasswordWhenIsInValid(t *testing.T) {
	adminBuilder := CreateAdminBuilder()
	adminBuilder.InitUser("gustavo", "12345")
	_ = adminBuilder.EncriptPassword()
	assert.False(t, adminBuilder.VerifyPassword("123456"))
}



