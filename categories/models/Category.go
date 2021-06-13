package models

type Item struct {
	Id string `json:"id"`
	Name string `json:"name"`
	RefId string `json:"ref_id"` //RefId is an id to reference to userId when it is category or categoryId when it's an subcategory
	IsAvailable bool `json:"is_available"`
	IsCategory bool `json:"is_category"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}