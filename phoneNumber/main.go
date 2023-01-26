package main

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

var (
	host   = "localhost"
	port   = 5432
	dbname = "gophercises_phone"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d database=%s sslmode=disable", host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	must(err)
	defer db.Close()
	must(db.Ping())

}

func must(err error) {
	if err != nil {
		fmt.Printf("Failed to call function \n err:%v", err)
		os.Exit(1)
	}
}

// func resetDB(db *sql.DB, name string) error {
// 	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
// 	if err != nil {
// 		return err
// 	}
// 	return createDB(db, name)
// }

// func createDB(db *sql.DB, name string) error {
// 	_, err := db.Exec("CREATE DATABASE " + name)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")

}
