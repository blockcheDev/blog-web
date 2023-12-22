package util

import "webback/db"

func AuthUserAndPassword(name string, password string) *db.User {
	// fmt.Printf("Name: %s, Password: %s", u.Name, u.Password)

	dbUser := &db.User{}
	res := db.DB.Where("Name = ? and Password = ?", name, password).First(dbUser)
	if res.RowsAffected == 0 {
		return nil
	}

	// c.SetCookie("ID", string(user.ID), -1, "/")
	return dbUser
}
