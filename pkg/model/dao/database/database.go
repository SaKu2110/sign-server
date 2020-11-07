package database

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
	db.UserDB.Close()
}
