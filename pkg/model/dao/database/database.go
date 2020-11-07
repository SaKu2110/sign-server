package database

import "github.com/SaKu2110/sign-server/pkg/model/service/log"

type DB struct {
	UserDB UserRepository
}

func New() (*DB, error) {
	userDB, err := newUserClient()
	if err != nil {
		return nil, err
	}
	return &DB{UserDB: userDB}, nil
}

func (db *DB) Close() {
	log.Info("Close Databases connection.")
	db.UserDB.Close()
}
