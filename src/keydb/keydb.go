package keydb

import (
	"github.com/robaho/leveldb"
)

type db struct {
	conn *leveldb.Database
}

// OpenDB initializes the database connection
func OpenDB(path string) (*db, error) {
	conn, err := leveldb.Open(path, leveldb.Options{})
	if err != nil {
		return nil, err
	}
	return &db{conn: conn}, nil
}

// CloseDB closes the database connection
func (db *db) CloseDB() error {
	return db.conn.Close()
}

// Get retrieves a value for a given key
func (db *db) Get(key string) (interface{}, error) {
	value, err := db.conn.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return string(value), nil
}

// Set the value of a key
func (db *db) Set(key string, value interface{}) error {
	return db.conn.Put([]byte(key), []byte(value.(string)))
}

// Delete a key
func (db *db) Delete(key string) error {
	_, err := db.conn.Remove([]byte(key))
	return err
}
