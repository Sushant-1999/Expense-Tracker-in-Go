package models

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories struct {
	CategoryList []Category `json:"categories"`
}
