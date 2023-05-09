package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Phone_number uint   `json:"nohp" form:"nohp"`
	Password     string `json:"password" form:"password"`
}

type UsersResponse struct {
	ID           uint   `json:"id" from:"id"`
	Name         string `json:"name" form:"name"`
	Phone_number uint   `json:"nohp" form:"nohp"`
	Token        string `json:"token" form:"token"`
}
