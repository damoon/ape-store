package model

type Product struct {
	Upc string `json:"upc"`
}

func (Product) IsEntity() {}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (User) IsEntity() {}
