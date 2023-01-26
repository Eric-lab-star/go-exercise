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
	createPhoneNumbersTable(db)
	id, err := insertPhonNumbers(db)
	must(err)
	fmt.Println("id: ", id)
}

func must(err error) {
	if err != nil {
		fmt.Printf("Failed to call function \n err:%v", err)
		os.Exit(1)
	}
}

func insertPhonNumbers(db *sql.DB) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	row := db.QueryRow(statement, "1234567890")
	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil

}

func createPhoneNumbersTable(db *sql.DB) {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers(
			id SERIAL,
			value VARCHAR(255)
		)
	`
	_, err := db.Exec(statement)
	must(err)
}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")

}
