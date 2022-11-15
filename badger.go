package util

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

type DB struct {
	DB    *badger.DB
	Error error
	Key   []byte
	Val   []byte
}

func (db DB) New(data string) *DB {
	db.DB, db.Error = badger.Open(badger.DefaultOptions(data))
	return &db
}

func (db *DB) Pull(key []byte) *DB {
	var item *badger.Item
	txn := db.DB.NewTransaction(false)
	defer txn.Discard()
	item, db.Error = txn.Get(key)
	if db.Error == nil {
		db.Val, db.Error = item.ValueCopy(db.Val)
	}
	return db
}

func (db *DB) Push(key, val []byte) *DB {
	txn := db.DB.NewTransaction(true)
	defer txn.Discard()
	db.Error = txn.Set(key, val)
	if db.Error == nil {
		db.Error = txn.Commit()
	}
	return db
}

func (db *DB) Drop(key []byte) *DB {
	txn := db.DB.NewTransaction(true)
	defer txn.Discard()
	db.Error = txn.Delete(key)
	if db.Error == nil {
		db.Error = txn.Commit()
	}
	return db
}

func (db *DB) Done() *DB {
	db.Error = db.DB.Close()
	return db
}

func (db *DB) Fatal() *DB {
	if db.Error != nil {
		log.Fatal(db.Error)
	}
	return db
}
