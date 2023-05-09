package database

import (
	models "myapp/model"
)

func GetNotifikasis() (interface{}, error) {
	var Notifikasis []models.Notifikasi

	err := DB.Find(&Notifikasis).Error
	if err != nil {
		return Notifikasis, err
	}

	return Notifikasis, nil
}

func GetNotifikasiByID(id int) (interface{}, error) {
	var Notifikasi models.Notifikasi

	err := DB.Where("id = ?", id).First(&Notifikasi).Error
	if err != nil {
		return Notifikasi, err
	}

	return Notifikasi, nil
}

func CreateNotifikasi(Notifikasi models.Notifikasi) (interface{}, error) {
	err := DB.Create(&Notifikasi).Error

	if err != nil {
		return Notifikasi, err
	}

	return Notifikasi, nil
}

func DeleteNotifikasi(id int) error {
	var Notifikasi models.Notifikasi

	result := DB.Where("id = ?", id).Delete(&Notifikasi)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateNotifikasi(id int, Notifikasi models.Notifikasi) error {
	result := DB.Model(&models.Notifikasi{}).Where("id = ?", id).Updates(Notifikasi)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
