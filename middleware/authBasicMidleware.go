package middleware

import (
	"mvc/config"
	"mvc/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	var user models.Users

	if err := config.DB.Where("email= ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
