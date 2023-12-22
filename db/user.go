package db

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string `gorm:"not null"`
	Password           string `gorm:"not null"`
	Email              string `gorm:"not null"`
	Telephone          string
	Gender             string
	jwt.StandardClaims `gorm:"-"`
}

func GetUser(ID uint) *User {
	user := User{}
	DB.Where("ID=?", ID).First(&user)
	return &user
}
func GetUserByName(name string) *User {
	user := User{}
	DB.Where("Name=?", name).First(&user)
	return &user
}
