package main

import (
	services "github.com/Julia-Marcal/reusable-api/config"
	db "github.com/Julia-Marcal/reusable-api/internal/database"
	router "github.com/Julia-Marcal/reusable-api/internal/http/router"
)

func main() {
	db.NewPostgres()
	services.NewDB()
	router.StartRouter()
}
