package util

import (
	"webback/db"
)

func AuthUserAndPassword(name string, password string) *db.User {
	dbUser := &db.User{}
	res := db.DB.Where("name = ?", name).First(dbUser)
	if res.RowsAffected == 0 {
		return nil
	}
	if !VerifyPassword(dbUser.Password, password) {
		return nil
	}

	return dbUser
}
