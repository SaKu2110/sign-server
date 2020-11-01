package service

import(
	"fmt"
	"log"
	"github.com/kelseyhightower/envconfig"
)

type dataBaseConfig struct {
	User string `envconfig:"DB_USER" default:"auth_user"`
	Pass string `envconfig:"DB_PASS" default:"password"`
	IP   string `envconfig:"DB_IP" default:"localhost"`
	Port string `envconfig:"DB_PORT" default:"3306"`
	Name string `envconfig:"DB_NAME" default:"auth"`
}

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func GetDataBaseConnectionToken() string {
	var c dataBaseConfig
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}
