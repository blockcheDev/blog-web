package main

import (
	"webback/db"
	"webback/router"
)

func main() {
	db.InitDatabase()
	router.StartRouter()
}
