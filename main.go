package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	controllers "github.com/Julia-Marcal/reusable-api/controllers"
	"github.com/Julia-Marcal/reusable-api/internal/db"
	internal "github.com/Julia-Marcal/reusable-api/internal/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	envGo()
	fmt.Println(os.Getenv("database"))

	dbcon, err := sql.Open("mysql", os.Getenv("database"))
	if err != nil {
		panic(err.Error())
	}
	defer dbcon.Close()

	dt := db.New(dbcon)
	ctx := context.Background()

	createUser := dt.CreateUser(ctx, internal.CreateUserParams{
		Name:      "Julia",
		Lastname:  "Marcal",
		Email:     "anajulia@gmail.com",
		Telephone: sql.NullString{String: "1194873942", Valid: true},
		Country:   sql.NullString{String: "Brazil", Valid: true},
	})
	if err != nil {
		panic(err.Error())
	}
	log.Println(createUser)

	router := gin.Default()
	router.GET("/ping", controllers.Pong)
	router.Run()
}
