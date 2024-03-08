package main

import (
	"webback/db"
	"webback/router"
)

func main() {
	db.InitDatabase()
	db.InitRedis()
	router.StartRouter()
}
