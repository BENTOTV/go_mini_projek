package database

import (
	utils "myapp/lib/middleware"
	models "myapp/model"
)

func GetAdmins() (interface{}, error) {
	var Admins []models.Admin

	err := DB.Find(&Admins).Error
	if err != nil {
		return Admins, err
	}

	return Admins, nil
}

func GetAdminByID(id int) (interface{}, error) {
	var Admin models.Admin

	err := DB.Where("id = ?", id).First(&Admin).Error
	if err != nil {
		return Admin, err
	}

	return Admin, nil
}

func CreateAdmin(Admin models.Admin) (interface{}, error) {
	hash, err := utils.HashPassword(Admin.Password)
	if err != nil {
		return Admin, err
	}

	Admin.Password = hash

	err = DB.Create(&Admin).Error
	if err != nil {
		return Admin, err
	}

	return Admin, nil
}

func DeleteAdmin(id int) error {
	var Admin models.Admin

	result := DB.Where("id = ?", id).Delete(&Admin)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateAdmin(id int, Admin models.Admin) error {
	result := DB.Model(&models.Admin{}).Where("id = ?", id).Updates(Admin)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func LoginAdmin(requestAdmin models.Admin) (models.Admin, error) {
	var Admin models.Admin

	err := DB.Where("nama = ?", requestAdmin.Nama).First(&Admin).Error
	if err != nil {
		return Admin, err
	}

	return Admin, nil
}
