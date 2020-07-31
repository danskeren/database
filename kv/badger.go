package kv

import (
	"sync"
	"time"

	"github.com/dgraph-io/badger/v2"
)

type badgerDB struct {
	badger *badger.DB
	mutex  sync.Mutex
}

func Open(opts badger.Options) (KV, error) {
	newDB, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &badgerDB{
		badger: newDB,
	}, err
}

func (db *badgerDB) Set(key, value []byte) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.badger.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (db *badgerDB) SetWithTTL(key, value []byte, duration time.Duration) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.badger.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(key, value).WithTTL(duration)
		return txn.SetEntry(entry)
	})
}

func (db *badgerDB) Get(key []byte) ([]byte, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	var value []byte
	err := db.badger.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}

func (db *badgerDB) Delete(key []byte) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.badger.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (db *badgerDB) GCTicker(discardRatio float64, duration time.Duration) {
	ticker := time.NewTicker(duration)
	stop := make(chan bool, 1)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := db.badger.RunValueLogGC(discardRatio)
			switch err {
			case badger.ErrRejected:
				stop <- true
			}
		case <-stop:
			return
		}
	}
}

func (db *badgerDB) Close() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if db.badger != nil {
		err := db.badger.Close()
		db.badger = nil
		return err
	}
	return nil
}
