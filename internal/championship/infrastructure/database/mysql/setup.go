package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DbName   string
	Username string
	Password string
	IP       string
	Port     string
	Protocol string
}

func GetConnection(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		c.Username,
		c.Password,
		c.Protocol,
		c.IP,
		c.Port,
		c.DbName,
	)
	return gorm.Open(mysql.Open(dsn))
}
