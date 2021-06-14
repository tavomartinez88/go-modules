//package for create users
package users

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/users/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const (
	RoleAdmin      		= "ADMIN"
	StatusAdmin         = "ACTIVE"
	FormatDateTimeAdmin = "01-02-2006 15:04:05"
)

// struct to represent to an admin
type adminBuilder struct{
	Id string
	UserName string
	Password string
	Role string
	Status string
	Created string
	Updated string
}

// newAdminBuilder create an admin instance
func newAdminBuilder() *adminBuilder {
	return &adminBuilder{}
}

//func to init an user, you provide username and password
func (a *adminBuilder) InitUser(username string, password string) {
	a.Id = uuid.New().String()
	a.UserName = username
	a.Password = password
}

// func to encript password
func (a *adminBuilder) EncriptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	} else {
  		a.Password = string(hash)
		return nil
	}
}

//func to verify if password plain provided match with password encripted
func (a *adminBuilder) VerifyPassword(plainPassword string) bool{
	bytePasswordHashed := []byte(a.Password)
	bytePasswordPlain := []byte(plainPassword)

	err:= bcrypt.CompareHashAndPassword(bytePasswordHashed, bytePasswordPlain)

	if err!= nil {
		log.Println(err)
		return false
	}

	return true
}

//func to set role to an admin
func (a *adminBuilder) SetRole() {
	a.Role = RoleAdmin
}

//func to get role
func (a *adminBuilder) GetRole() string {
	return a.Role
}

//func for set create date and update date
func (a *adminBuilder) SetDateTimeBuilding() {
	now := time.Now()
	a.Created = now.Format(FormatDateTimeAdmin)
	a.Updated = now.Format(FormatDateTimeAdmin)
}

//func to set username
func (a *adminBuilder) SetUsername(username string) {
	a.UserName = username
}

//func to set password
func (a *adminBuilder) SetPassword(password string) {
	a.Password = password
}

// func get user as admin
func (a *adminBuilder) GetUser() models.User{
	a.InitUser(a.UserName, a.Password)
	a.SetRole()
	a.SetStatus()
	a.SetDateTimeBuilding()
	_ = a.EncriptPassword()
	return models.User{
		Id: a.Id,
		UserName: a.UserName,
		Password: a.Password,
		Role: a.Role,
		Status: a.Status,
		Created: a.Created,
		UserNameVerified: false,
		Updated: a.Updated,
	}
}

//func to set status active to admin
func (a *adminBuilder) SetStatus() {
	a.Status = StatusAdmin
}
