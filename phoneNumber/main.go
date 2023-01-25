package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

var (
	host   = "localhost"
	port   = 5432
	dbname = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d sslmode=disable", host, port)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = createDB(db, dbname)
	if err != nil {
		panic(err)
	}
	db.Close()

}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")

}
