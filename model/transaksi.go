package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	UserID      uint `json:"user_id"`
	Tanggal     time.Time
	LayananID   uint `json:"layanan_id"`
	Jumlah      int  `json:"jumlah"`
	Total_harga int  `json:"total_harga"`
}
