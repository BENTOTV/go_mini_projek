package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Phone_number uint   `json:"nohp" form:"nohp"`
	Password     string `json:"password" form:"password"`
}
