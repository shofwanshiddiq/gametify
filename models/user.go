package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
user_id
email
password
fullname
*/

type RoleType string

const (
	TypeAdmin RoleType = "admin"
	TypeUser  RoleType = "user"
)

type User struct {
	gorm.Model
	Name           string   `json:"name"`
	Email          string   `json:"email" gorm:"unique"`
	Password       string   `json:"password"`
	ProfilePicture string   `json:"profile_picture"`
	Role           RoleType `json:"role" gorm:"type:varchar(20)"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
