package users

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/users/models"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const (
	RoleClient           = "CLIENT"
	FormatDateTimeClient = "01-02-2006 15:04:05"
)

type clientBuilder struct{
	Id string
	UserName string
	Password string
	Role string
	Address utils.Address
	Phone utils.Phone
	Created string
	Updated string
}

func newClientBuilder() *clientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) InitUser(username string, password string) {
	c.Id = uuid.New().String()
	c.UserName = username
	c.Password = password
}

func (c *clientBuilder) EncriptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	} else {
		c.Password = string(hash)
		return nil
	}
}

func (a *clientBuilder) SetUsername(username string) {
	a.UserName = username
}

func (a *clientBuilder) SetPassword(password string) {
	a.Password = password
}

func (c *clientBuilder) VerifyPassword(plainPassword string) bool {
	bytePasswordHashed := []byte(c.Password)
	bytePasswordPlain := []byte(plainPassword)

	err:= bcrypt.CompareHashAndPassword(bytePasswordHashed, bytePasswordPlain)

	if err!= nil {
		log.Println(err)
		return false
	}

	return true
}

func (c *clientBuilder) SetRole() {
	c.Role = RoleClient
}

func (c *clientBuilder) GetRole() string {
	return c.Role
}

func (c *clientBuilder) SetDateTimeBuilding() {
	now := time.Now()
	c.Created = now.Format(FormatDateTimeClient)
	c.Updated = now.Format(FormatDateTimeClient)
}

//methods concretes

func (c *clientBuilder) GetUser() models.User{
	c.InitUser(c.UserName, c.Password)
	c.SetRole()
	c.SetDateTimeBuilding()
	_ = c.EncriptPassword()
	user := models.User{
		Id: c.Id,
		UserName: c.UserName,
		Password: c.Password,
		Role: c.Role,
		Created: c.Created,
		UserNameVerified: false,
		Updated: c.Updated,
	}

	if c.Address.Id != "" {
		user.Addresses = []utils.Address{c.Address}
	}

	if c.Phone.Id != "" {
		user.Phones = []utils.Phone{c.Phone}
	}

	return user
}

func (c *clientBuilder) SetAddress(address utils.Address) {
	c.Address = address
}

func (c *clientBuilder) SetPhone(phone utils.Phone) {
	c.Phone = phone
}