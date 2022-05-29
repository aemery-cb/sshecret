package secret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

type SecretManager struct {
	db *badger.DB
}

func (sm *SecretManager) Clean() {
	sm.db.Close()
}

func (sm *SecretManager) Get(key string) (string, error) {
	var valCopy []byte = nil
	err := sm.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	return string(valCopy), err
}
func (sm *SecretManager) Put(key, value string) error {
	err := sm.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	return err
}
func (sm *SecretManager) Delete() {

}
func NewSecretManager() *SecretManager {
	db, err := badger.Open(badger.DefaultOptions("badger.db"))
	if err != nil {
		log.Fatal(err)
	}

	return &SecretManager{db: db}
}

func hotp(key string, counter uint64, digits int) int {
	bytes := []byte(key)
	h := hmac.New(sha1.New, bytes)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func totp(key string, t time.Time, digits int) int {
	return hotp(key, uint64(t.UnixNano())/30e9, digits)
}
