package controllers

import (
	"mvc/lib/database"
	ownMid "mvc/middleware"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(echoConteks echo.Context) error {
	user, err := database.GetAllUsers()
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoConteks.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses",
		"data":   newResponesArr(*user),
	})
}

func GetUserControllerFilter(echoConteks echo.Context) error {

	name := echoConteks.QueryParam("name")
	user, err := database.GetAllUsersFilter(name, 4, 1)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoConteks.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses",
		"data":   newResponesArr(*user),
	})
}

func GetUserbyId(echoConteks echo.Context) error {
	id, err := strconv.Atoi(echoConteks.Param("id"))

	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	user, err := database.GetUserByID(id)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoConteks.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses",
		"data":   newRespones(*user),
	})
}

func SaveUserController(echoConteks echo.Context) error {
	var users RequestUser

	echoConteks.Bind(&users)

	result, err := database.StoreUser(users.toModel())

	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoConteks.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses",
		"data":   newRespones(*result),
	})
}

func DeletedUserByID(echoConteks echo.Context) error {
	var user models.Users
	id, _ := strconv.Atoi(echoConteks.Param("id"))
	_, errFind := database.GetUserByID(id)
	_, err := database.DeletedUserByID(id, user)
	if errFind != nil {
		return echoConteks.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return echoConteks.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses delete",
	})
}

func UpdateUsersController(echoConteks echo.Context) error {
	var updateUser RequestUser

	echoConteks.Bind(&updateUser)

	id, _ := strconv.Atoi(echoConteks.Param("id"))

	result, err := database.UpdateUser(id, updateUser.toModel())

	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return echoConteks.JSON(http.StatusOK, newRespones(*result))
}

func GetUserToken(echoContext echo.Context) error {

	username := echoContext.QueryParam("username")
	password := echoContext.QueryParam("password")

	token := ""
	var err error
	if username == "kamil" && password == "masukaja" {
		token, err = ownMid.CreateToken(1)
	}

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"data":   err,
		})
	} else {
		return echoContext.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   token,
		})
	}
}

func HelloWorld(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusAccepted, "Logged in")
}
