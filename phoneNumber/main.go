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

/*
numbers in phone_numbers table
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
*/

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d database=%s sslmode=disable", host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	must(err)
	defer db.Close()
	createPhoneNumbersTable(db)
	phoneNumbers, err := getAllPhoneNumbers(db)

	must(err)

	for _, phoneNumber := range phoneNumbers {

		fmt.Printf("id: %d, phone number: %s\n", phoneNumber.id, normalize(phoneNumber.value))
	}

}

func must(err error) {
	if err != nil {
		fmt.Printf("Failed to call function \n err:%v\n", err)
		os.Exit(1)
	}
}

type phoneNumber struct {
	id    int
	value string
}

func getAllPhoneNumbers(db *sql.DB) ([]phoneNumber, error) {
	var ret []phoneNumber
	statement := `SELECT * FROM phone_numbers`
	rows, err := db.Query(statement)
	if err != nil {
		return ret, nil
	}
	for rows.Next() {
		var phoneNumber phoneNumber
		err := rows.Scan(&phoneNumber.id, &phoneNumber.value)
		if err != nil {
			return nil, err
		}
		ret = append(ret, phoneNumber)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func getPhoneNumbers(db *sql.DB, id int) (string, error) {
	var value string
	statement := `SELECT value FROM phone_numbers WHERE id=$1`
	err := db.QueryRow(statement, id).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, err

}

// func insertPhonNumbers(db *sql.DB, phoneNumber string) (int, error) {
// 	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
// 	row := db.QueryRow(statement, phoneNumber)
// 	var id int
// 	err := row.Scan(&id)
// 	if err != nil {
// 		return -1, err
// 	}

// 	return id, nil

// }

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
