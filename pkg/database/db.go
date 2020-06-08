package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Config struct {
	Name, Host     string
	Port           int
	User, Password string
	Debug          bool
}

func New(config Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.Name))
	if err != nil {
		return nil, err
	}
	if config.Debug {
		db = db.Debug()
	}
	db.LogMode(config.Debug)
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db, nil
}
