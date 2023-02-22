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
	Host     string
	Port     string
	Protocol string
}

func GetConnection(c Config) (*gorm.DB, error) {
	port := "3306"
	if c.Port != "" {
		port = c.Port
	}

	protocol := "tcp"
	if c.Protocol != "" {
		protocol = c.Protocol
	}

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		c.Username,
		c.Password,
		protocol,
		c.Host,
		port,
		c.DbName,
	)
	return gorm.Open(mysql.Open(dsn))
}
