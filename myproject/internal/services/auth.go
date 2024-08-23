package services

import (
	"myproject/internal/models"
	"myproject/pkg/database"
	"myproject/pkg/utils"
)

func RegisterUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?,?)", username, hashedPassword)
	return err
}

func LoginUser(username, password string) (bool, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT * FROM users WHERE username =?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return false, err
	}
	return utils.CheckPasswordHash(password, user.Password), nil
}
