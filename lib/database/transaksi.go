package database

import (
	models "myapp/model"
)

func GetTransaksis() (interface{}, error) {
	var Transaksis []models.Transaksi

	err := DB.Find(&Transaksis).Error
	if err != nil {
		return Transaksis, err
	}

	return Transaksis, nil
}

func GetTransaksiByID(id int) (interface{}, error) {
	var Transaksi models.Transaksi

	err := DB.Where("id = ?", id).First(&Transaksi).Error
	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

func CreateTransaksi(Transaksi models.Transaksi) (interface{}, error) {
	err := DB.Create(&Transaksi).Error

	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

func DeleteTransaksi(id int) error {
	var Transaksi models.Transaksi

	result := DB.Where("id = ?", id).Delete(&Transaksi)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateTransaksi(id int, Transaksi models.Transaksi) error {
	result := DB.Model(&models.Transaksi{}).Where("id = ?", id).Updates(Transaksi)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
