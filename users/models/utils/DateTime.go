package utils

type Day struct {
	Id string `json:"id"`
	Name string `json:"name" validate:"required, min=5, max=9"`
	Time Time `json:"time" validate:"required"`
}

type Time struct {
	Time TimeFromTo `json:"time" validate:"required"`
	Breaks [] TimeFromTo `json:"breaks"`
}

type TimeFromTo struct {
	From string `json:"from" validate:"required, min=3, max=5"`
	To string `json:"to" validate:"required, min=3, max=5"`
}
