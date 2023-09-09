package main

import (
	db "github.com/Julia-Marcal/reusable-api/repository/database"
	router "github.com/Julia-Marcal/reusable-api/router"
)

func main() {
	db.NewPostgres()
	router.StartRouter()
}
