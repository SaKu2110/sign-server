package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type userDBAccessConfig struct {
	User string `envconfig:"DB_USER" default:"auth_user"`
	Pass string `envconfig:"DB_PASS" default:"password"`
	IP   string `envconfig:"DB_IP" default:"localhost"`
	Port string `envconfig:"DB_PORT" default:"3306"`
	Name string `envconfig:"DB_NAME" default:"auth"`
}

type dbConnConfig struct {
	MaxOpenConns int `envconfig:"DB_MAX_CONNS" default:"100"`
	MaxIdle      int `envconfig:"DB_MAX_IDLE" default:"100"`
	MaxLifeTime  int `envconfig:"DB_MAX_LIFETIME" default:"100"`
}

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func UserDBAccessToken() (string, error) {
	var c userDBAccessConfig
	if err := envconfig.Process("", &c); err != nil {
		return "", err
	}

	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name), nil
}

func DBConnectionInfo() (int, int, time.Duration, error) {
	var c dbConnConfig
	if err := envconfig.Process("", &c); err != nil {
		return 0, 0, 0, err
	}

	return c.MaxOpenConns, c.MaxIdle, time.Duration(c.MaxLifeTime) * time.Second, nil
}
