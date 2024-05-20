package models

import "gorm.io/gorm"

type Employees struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Line_ID  string `json:"line_id"`
	Tel      string `json:"tel" validate:"required,number,min=10,max=10"`
	Business string `json:"business" validate:"required"`
	Website  string `json:"website" validate:"required,min=2,max=30"`
}

type Dogs struct {
	gorm.Model        // ID , CreateAT , IpdateAT ,DeleteAT
	Name       string `json:"name"`
	DogID      int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Count       int       `json:"count"`
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Sum_Red     int       `json:"sum_red"`
	Sum_Green   int       `json:"sum_green"`
	Sum_Pink    int       `json:"sum_pink"`
	Sum_Nocolor int       `json:"nocolor"`
}

type Company struct {
	gorm.Model            // ID , CreateAT , IpdateAT ,DeleteAT
	CompanyName    string `json:"company_name"`
	CompanyType    string `json:"company_type"`
	CompanyAddress string `json:"company_address"`
	CompanyPeople  int    `json:"company_people"`
	CompanyTel     string `json:"company_tel"`
	CompanyID      int    `json:"company_id"`
}
