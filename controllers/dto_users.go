package controllers

import (
	"mvc/models"
)

type RequestUser struct {
	Nik      int
	Name     string
	Alamat   string
	Email    string
	Password string
}

type ResponseUser struct {
	ID       int    `json:"id"`
	Nik      int    `json:"nik"`
	Name     string `json:"name" `
	Alamat   string `json:"alamat"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (re *RequestUser) toModel() models.Users {
	return models.Users{
		Nik:      re.Nik,
		Name:     re.Name,
		Alamat:   re.Alamat,
		Email:    re.Email,
		Password: re.Password,
	}
}

func newRespones(users models.Users) ResponseUser {
	return ResponseUser{
		ID:       int(users.ID),
		Nik:      users.Nik,
		Name:     users.Name,
		Alamat:   users.Alamat,
		Email:    users.Email,
		Password: users.Password,
	}
}
func newResponesArr(users []models.Users) []ResponseUser {
	var resp []ResponseUser
	for _, value := range users {
		resp = append(resp, newRespones(value))
	}
	return resp
}
