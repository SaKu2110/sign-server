package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/SaKu2110/sign-server/pkg/model/service/log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/SaKu2110/sign-server/pkg/model/dto"

	"github.com/SaKu2110/sign-server/pkg/model/service/config"
)

const (
	BACKOFF_MAX_CONNS          = 5
	QUERY_FORMAT_GET_USER_INFO = "SELECT * FROM `user` WHERE id = ?"
	QUERY_FORMAT_ADD_USER_INFO = "INSERT INTO `user` (id, password) VALUES (?, ?)"
)

type UserRepository interface {
	Get(string) ([]dto.UserInfo, error)
	Add(string, string) error
	Close()
}
type userRepository struct {
	DB *sql.DB
}

func newUserClient() (UserRepository, error) {
	var db *sql.DB
	var err error
	token, err := config.UserDBAccessToken()
	if err != nil {
		return nil, err
	}
	for i := 0; ; i++ {
		if i == BACKOFF_MAX_CONNS {
			return nil, errors.New("Faild connection user database")
		}
		db, err = sql.Open("mysql", token)
		if err != nil {
			return nil, err
		}
		// check connection database //
		if err := db.Ping(); err == nil {
			break
		}
		log.Warn(fmt.Sprintf("Faild connection `user` Database. (retry: %d)", i))
		// Exponential Backoff //
		backoff(i)
	}
	if err := setConnConfigs(db); err != nil {
		return nil, err
	}
	return &userRepository{DB: db}, nil
}
func setConnConfigs(db *sql.DB) error {
	maxConn, maxIdle, maxlifetime, err := config.DBConnectionInfo()
	log.Info(fmt.Sprintf("DB Connection info: Max open connections->%d", maxConn))
	log.Info(fmt.Sprintf("DB Connection info: Max idle->%d", maxIdle))
	log.Info(fmt.Sprintf("DB Connection info: Max lifetime->%ds", maxlifetime))
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(maxlifetime)
	return nil
}

func backoff(i int) {
	time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
}
func (r *userRepository) Get(id string) ([]dto.UserInfo, error) {
	var users []dto.UserInfo
	rows, err := r.DB.Query(QUERY_FORMAT_GET_USER_INFO, id)
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
	return users, nil
}

func (r *userRepository) Add(id, password string) error {
	result, err := r.DB.Prepare(QUERY_FORMAT_ADD_USER_INFO)
	if err != nil {
		return err
	}
	_, err = result.Exec(id, password)
	return err
}

func (r *userRepository) Close() {
	r.DB.Close()
}
