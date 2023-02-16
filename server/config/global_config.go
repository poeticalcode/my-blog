package config

import (
	"fmt"
)

type mysqlConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DataBase string
	Charset  string
}

type githubOOS struct {
	Token      string
	Message    string
	Repository string
	API        string
	UserName   string
}

func (db mysqlConfig) Dsn() (dsn string) {
	//"utf8mb4"
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		db.UserName, db.Password, db.Host, db.Port, db.DataBase, db.Charset)
	return
}

type globalConfig struct {
	ProjectRootPath string
	Port            string
	MySQL           mysqlConfig
	GithubOOS       githubOOS
}
