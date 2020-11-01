package dao

import(
	"fmt"
	"time"
	"math"

	// import gorm library
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/SaKu2110/sign-server/pkg/model/dto"
	"github.com/SaKu2110/sign-server/pkg/model/service"
)

const(
	QUERY_FORMAT_GET_USER_INFO = "SELECT * FROM `user` WHERE id = ?"
	QUERY_FORMAT_SET_USER_INFO = "INSERT INTO `user` (id, password) VALUES (?, ?)"
)

type DB struct {
	DB	*sql.DB
}

func Init() (*DB, error) {
	token := service.GetDataBaseConnectionToken()
	for i:=0; i<5; i++ {
		db, err := sql.Open("mysql", token)
		if err != nil {
			return nil, err
		}
		// check connection database //
		if err := db.Ping(); err == nil {
			return &DB{DB: db}, nil
		}
		// Exponential Backoff //
		backoff(i)
	}
	return nil, fmt.Errorf("Faild connection database")
}

func backoff(i int) {
	time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
}

func (db *DB) Close() {
	db.DB.Close()
}

func (db *DB) GetUserInfo(id string) (users []dto.UserInfo, err error) {
	rows, err := db.DB.Query(QUERY_FORMAT_GET_USER_INFO, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, pass string
		if err := rows.Scan(&id, &pass); err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		users = append(users, dto.UserInfo{ID: id, Password: pass})
	}
	return
}

func (db *DB) InsertUserInfo(id string, password string) error {
	result, err := db.DB.Prepare(QUERY_FORMAT_SET_USER_INFO)
	if err != nil {
		return err
	}
	_, err = result.Exec(id, password)
	return err
}

