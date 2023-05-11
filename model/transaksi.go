package model

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	UserID      uint `json:"user_id"`
	LayananID   uint `json:"layanan_id"`
	Jumlah      int  `json:"jumlah"`
	Total_harga int  `json:"total_harga"`
}
