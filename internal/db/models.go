// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
)

type User struct {
	ID        int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      string
	Lastname  string
	Email     string
	Telephone sql.NullString
	Country   sql.NullString
}
