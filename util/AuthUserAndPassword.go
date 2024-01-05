package util

import "webback/db"

func AuthUserAndPassword(name string, password string) *db.User {
	dbUser := &db.User{}
	res := db.DB.Where("name = ? and password = ?", name, password).First(dbUser)
	if res.RowsAffected == 0 {
		return nil
	}

	return dbUser
}
