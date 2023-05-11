package database

import (
	models "myapp/model"
)

func GetLayanans() (interface{}, error) {
	var Layanans []models.Layanan

	err := DB.Find(&Layanans).Error
	if err != nil {
		return Layanans, err
	}

	return Layanans, nil
}

func GetLayananByID(id int) (interface{}, error) {
	var Layanan models.Layanan

	err := DB.Where("id = ?", id).First(&Layanan).Error
	if err != nil {
		return Layanan, err
	}

	return Layanan, nil
}

func CreateLayanan(Layanan models.Layanan) (interface{}, error) {

	err := DB.Create(&Layanan).Error

	if err != nil {
		return Layanan, err
	}

	return Layanan, nil
}

func DeleteLayanan(id int) error {
	var Layanan models.Layanan

	result := DB.Where("id = ?", id).Delete(&Layanan)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateLayanan(id int, Layanan models.Layanan) error {
	result := DB.Model(&models.Layanan{}).Where("id = ?", id).Updates(Layanan)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
