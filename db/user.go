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
	Gender             string `gorm:"not null"`
	jwt.StandardClaims `gorm:"-"`
}

func GetUser(ID any) *User {
	user := User{}
	DB.Where("id=?", ID).First(&user)
	return &user
}
func GetUserByName(name string) *User {
	user := User{}
	DB.Where("name=?", name).First(&user)
	return &user
}
