package models

type Contact struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Numbers   []int  `json:"numbers"`
}

