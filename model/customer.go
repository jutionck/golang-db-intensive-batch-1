package model

import "database/sql"

type Customer struct {
	Id      int
	Name    string
	Address sql.NullString
	Job     string
}
