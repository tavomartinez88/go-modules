package users

import (
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/users/models"
	"github.com/tavomartinez88/go-modules/users/models/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const RoleProvider = "PROVIDER"
const StatusProvider = "ACTIVE"

type providerBuilder struct{
	Id string
	UserName string
	Password string
	Role string
	Addresses []utils.Address
	WebSite string
	FacebookProfile string
	InstagramProfile string
	TwitterProfile string
	Phones []utils.Phone
	Status string
	AttentionDays []utils.Day
	Created string
	Updated string
}

func newProviderBuilder() *providerBuilder {
	return &providerBuilder{}
}

func (p *providerBuilder) InitUser(username string, password string) {
	p.Id = uuid.New().String()
	p.UserName = username
	p.Password = password
}

func (p *providerBuilder) SetUsername(username string) {
	p.UserName = username
}

func (p *providerBuilder) SetPassword(password string) {
	p.Password = password
}

func (p *providerBuilder) EncriptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	} else {
		p.Password = string(hash)
		return nil
	}
}

func (p *providerBuilder) VerifyPassword(plainPassword string) bool {
	bytePasswordHashed := []byte(p.Password)
	bytePasswordPlain := []byte(plainPassword)

	err:= bcrypt.CompareHashAndPassword(bytePasswordHashed, bytePasswordPlain)

	if err!= nil {
		log.Println(err)
		return false
	}

	return true
}

func (p *providerBuilder) SetRole() {
	p.Role = RoleProvider
}

func (p *providerBuilder) GetRole() string {
	return p.Role
}

func (p *providerBuilder) SetDateTimeBuilding() {
	now := time.Now()
	p.Created = now.Format(FormatDateTimeAdmin)
	p.Updated = now.Format(FormatDateTimeAdmin)
}

// methods concretes
func (p *providerBuilder) SetStatus() {
	p.Status = StatusProvider
}

func (p *providerBuilder) GetUser() models.User{
	p.InitUser(p.UserName, p.Password)
	p.SetRole()
	p.SetDateTimeBuilding()
	_ = p.EncriptPassword()
	return models.User{
		Id: p.Id,
		UserName: p.UserName,
		Password: p.Password,
		Role: p.Role,
		Addresses: p.Addresses,
		WebSite: p.WebSite,
		FacebookProfile: p.FacebookProfile,
		InstagramProfile: p.InstagramProfile,
		TwitterProfile: p.TwitterProfile,
		Phones: p.Phones,
		Status: p.Status,
		AttentionDays: p.AttentionDays,
		Created: p.Created,
		Updated: p.Updated,
	}
}

func (p *providerBuilder) SetAddresses(addresses []utils.Address) {
	p.Addresses = addresses
}

func (p *providerBuilder) SetPhones(phones []utils.Phone) {
	p.Phones = phones
}

func (p *providerBuilder) SetWebSite(webSite string) {
	p.WebSite = webSite
}

func (p *providerBuilder) SetFacebookProfile(facebook string) {
	p.FacebookProfile = facebook
}

func (p *providerBuilder) SetTwitterProfile(twitter string) {
	p.TwitterProfile = twitter
}

func (p *providerBuilder) SetInstagramProfile(instagram string) {
	p.InstagramProfile = instagram
}

func (p *providerBuilder) SetAttentionDays(days [] utils.Day) {
	p.AttentionDays = days
}