package model

import (
	"gorm.io/gorm"
)

type Notifikasi struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	TransaksiID uint   `json:"transaksi_id"`
	Status      string `json:"status" form:"status"`
	Transaksi   Transaksi
	User        User
}
