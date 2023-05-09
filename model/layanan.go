package model

import (
	"gorm.io/gorm"
)

type Layanan struct {
	gorm.Model
	Nama  string `json:"nama" form:"nama"`
	Harga int    `json:"harga" form:"harga"`
	Type  string `json:"type" form:"type"`
}
