package models

type Category struct {
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
	Color string `json:"color"`
}
