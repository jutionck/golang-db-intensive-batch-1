package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// Langkah cara untuk melakukan koneksi ke database
	// 1.github.com/lib/pq -> merupakan sebuah driver postgre (connector dan perlu di install)
	// 2.database/sql (built-in go) -> ini untuk query -> SELECT, INSERT, UPDATE, DELETE
	// 3. (opsi) -> jmoiron/sqlx

	// Kita siapkan koneksi
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"
	dbName := "db_enigma"
	dbDriver := "postgres"

	// cara koneksi ke db itu ada dua (2)
	// 1. url -> "dbDriver://dbUser:dbPassword@dbHost:dbPort/dbName"
	// 2. standard -> "dbuser=xx dbname=xx dbpassword=xx dbport=xx"
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		fmt.Println(err)
	}

	result, err := db.Exec("insert into customer(name,address,job) values ('Bulan', 'Bali', 'Penyanyi')")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
