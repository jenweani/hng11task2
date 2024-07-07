package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct{
	UserId string `json:"userId" gorm:"primaryKey;type:varchar(255)"`
	// Add user fields

	FirstName    string `gorm:"column:firstName;not null" json:"firstName"`
	LastName     string `gorm:"column:lastName;not null" json:"lastName"`
	Email        string `gorm:"column:email;unique;not null" json:"email"`
	Password	string `gorm:"column:password;not null" json:"password"`
	Phone  string `gorm:"column:phone" json:"phone"`
	Organisations []Organisation 
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UserId = uuid.NewString()
	password, err := HashPassword(u.Password) //Hash password before creating the user
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) ResponseMap() map[string]interface{}{
	return map[string]interface{}{
		"userId": u.UserId,
		"firstName": u.FirstName,
		"lastName": u.LastName,
		"email": u.Email,
		"phone": u.Phone,
	}
}