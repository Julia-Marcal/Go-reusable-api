package main

import (
	router "github.com/Julia-Marcal/reusable-api/router"
	newdb "github.com/Julia-Marcal/reusable-api/services"
)

func main() {
	newdb.NewDB()
	router.Start_Router()
}