package config

import "fmt"

type mysqlConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DataBase string
	Charset  string
}

func (db mysqlConfig) Dsn() (dsn string) {
	//"utf8mb4"
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		db.UserName, db.Password, db.Host, db.Port, db.DataBase, db.Charset)
	return
}

type globalConfig struct {
	Port  string
	MySQL mysqlConfig
}
