package db

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string `gorm:"not null; index:idx_name_password,priority:1"`
	Password           string `gorm:"not null; index:idx_name_password,priority:2"`
	Email              string `gorm:"not null"`
	Telephone          string
	Gender             string `gorm:"not null"`
	AvaterUrl          string
	IsAdmin            bool `gorm:"DEFAULT false"`
	jwt.StandardClaims `gorm:"-"`
}

func GetUser(ID any) *User {
	user := User{}
	res := DB.Where("id=?", ID).First(&user)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}
func GetUserByName(name string) *User {
	user := User{}
	res := DB.Where("name=?", name).First(&user)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}
func (user *User) VerifyAdmin() bool {
	return user.IsAdmin
}
