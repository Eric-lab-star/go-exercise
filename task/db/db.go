package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("task")

type Task struct {
	ID    int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err

	})
	return err

}
