package util

import (
	"webback/db"
)

// 注意！！！函数返回的结构体里只有name, password字段是有值的
func AuthUserAndPassword(name string, password string) *db.User {
	dbUser := &db.User{}
	res := db.DB.Select("name", "password").Where("name = ?", name).First(dbUser)
	if res.RowsAffected == 0 {
		return nil
	}
	if !VerifyPassword(dbUser.Password, password) {
		return nil
	}

	return dbUser
}
