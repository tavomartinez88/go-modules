package models

import "github.com/tavomartinez88/go-modules/users/models/utils"

type User struct {
	Id string `json:"id" validate:"required"`
	UserName string `json:"user_name"`
	UserNameVerified bool `json:"user_name_verified"`
	Password string `json:"password" validate:"required"`
	Role string `json:"role" validate:"required,alpha"`
	Addresses []utils.Address `json:"addresses"`
	WebSite string `json:"web_site"`
	FacebookProfile string `json:"facebook_profile"`
	InstagramProfile string `json:"instagram_profile"`
	TwitterProfile string `json:"twitter_profile"`
	Phones []utils.Phone `json:"phones"`
	Status string `json:"status" validate:"required,alpha"`
	AttentionDays []utils.Day `json:"attention_days"`
	Created string `json:"created" validate:"required"`
	Updated string `json:"updated" validate:"required"`
}
