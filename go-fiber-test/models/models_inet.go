package models

type Employees struct {
	Email    string `json:"email" validate:"required,email,min=3,max=32"`
	Username string `json:"username" validate:"required,uppercase,lowercase,number"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Line_ID  string `json:"line_id"`
	Tel      string `json:"tel" validate:"required"`
	Business string `json:"business" validate:"required"`
	Website  string `json:"website" validate:"required"`
}
