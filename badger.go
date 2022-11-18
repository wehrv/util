package util

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

var dB *badger.DB

type DB struct {
	DB    *badger.DB
	Error error
	Item  *badger.Item
	Txn   *badger.Txn
	Key   []byte
	Val   []byte
}

func Init() {
	var err error
	dB, err = badger.Open(badger.DefaultOptions("data"))
	if err != nil {
		log.Fatal(err)
	}
}

func (db DB) Init(tf bool) *DB {
	db.DB = dB
	db.Txn = db.DB.NewTransaction(tf)
	return &db
}

func (db *DB) Done() *DB {
	if db.Error == nil {
		db.Error = db.Txn.Commit()
		if db.Error == nil {
			db.Error = db.DB.Close()
		}
	}
	db.Txn.Discard()
	return db
}

func (db *DB) Pull(key []byte) *DB {
	db.Key = key
	db.Item, db.Error = db.Txn.Get(key)
	if db.Error == nil {
		db.Val, db.Error = db.Item.ValueCopy(db.Val)
	}
	return db
}

func (db *DB) Push(key, val []byte) *DB {
	db.Key, db.Val = key, val
	db.Error = db.Txn.Set(key, val)
	if db.Error == nil {
		db.Error = db.Txn.Commit()
	}
	return db
}

func (db *DB) Drop(key []byte) *DB {
	db.Key = key
	db.Error = db.Txn.Delete(key)
	if db.Error == nil {
		db.Error = db.Txn.Commit()
	}
	return db
}
