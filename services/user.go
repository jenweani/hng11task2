package services

import (
	"errors"
	"hng11task2/internal/db"
	"hng11task2/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsersByEmail(email string) []models.User {
	var users []models.User

	db.DB.Where("email = ?", email).Find(&users)
	return users
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateNewUser(firstName, lastName, email, password, phone string) (*models.User, error) {
	orgName := firstName + "'s Organisation"
	err := db.DB.Create(&models.User{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		Phone: phone,
		Organisations: []models.Organisation{
			{Name: orgName, Description: "New Organisation"},
		},

	}).Error
	if err != nil {
		return nil, err
	}
	var user models.User
	err = db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user,nil
}

func GetUserById(userId string) (*models.User, error) {
	var user models.User

	result := db.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, result.Error 
	}

	return &user, nil
}

func GetUsersByEmailAndPassword(email, password string) (*models.User, error) {
	var user models.User

	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, result.Error 
	}

	match := VerifyPassword(password, user.Password)
	if !match {
		return nil, errors.New("password does not match")
	}
	return &user, nil
}