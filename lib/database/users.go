package database

import (
	"mvc/config"
	"mvc/models"
)

func GetAllUsers() (*[]models.Users, error) {
	var user []models.Users
	if err := config.DB.Find(&user).Error; err != nil {
		return &[]models.Users{}, err
	}
	return &user, nil
}
func GetAllUsersFilter(name string, rowPerpage int, page int) (*[]models.Users, error) {
	var user []models.Users

	if rowPerpage == 0 {
		rowPerpage = 5
	}
	if page == 0 {
		page = 1
	}
	query := config.DB

	if name != "" {
		query = query.Where("name LIKE '%?%'", name)
	}
	query.Limit(rowPerpage)
	query.Offset((rowPerpage * page) - rowPerpage)
	if err := config.DB.Find(&user).Error; err != nil {
		return &[]models.Users{}, err
	}
	return &user, nil
}

func GetUserByID(id int) (*models.Users, error) {
	var user models.Users
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return &models.Users{}, err
	}
	return &user, nil
}

func StoreUser(user models.Users) (*models.Users, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return &models.Users{}, err
	}
	return &user, nil
}

func DeletedUserByID(id int, user models.Users) (*models.Users, error) {
	if err := config.DB.Unscoped().Delete(&user, id).Error; err != nil {
		return &models.Users{}, err
	}
	return &user, nil
}

func UpdateUser(id int, user models.Users) (*models.Users, error) {
	var users models.Users
	if err := config.DB.Find(&users, id).Updates(&user).Error; err != nil {
		return &models.Users{}, err
	}
	return &user, nil
}
