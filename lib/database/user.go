package database

import (
	utils "myapp/lib/middleware"
	models "myapp/model"
)

func GetUsers() (interface{}, error) {
	var Users []models.User

	err := DB.Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}

func GetUserByID(id int) (interface{}, error) {
	var User models.User

	err := DB.Where("id = ?", id).First(&User).Error
	if err != nil {
		return User, err
	}

	return User, nil
}

func CreateUser(User models.User) (interface{}, error) {
	hash, err := utils.HashPassword(User.Password)
	if err != nil {
		return User, err
	}

	User.Password = hash

	err = DB.Create(&User).Error
	if err != nil {
		return User, err
	}

	return User, nil
}

func DeleteUser(id int) error {
	var User models.User

	result := DB.Where("id = ?", id).Delete(&User)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateUser(id int, User models.User) error {
	result := DB.Model(&models.User{}).Where("id = ?", id).Updates(User)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func LoginUser(requestUser models.User) (models.User, error) {
	var User models.User

	err := DB.Where("nama = ?", requestUser.Nama).First(&User).Error
	if err != nil {
		return User, err
	}

	return User, nil
}
