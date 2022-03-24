package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	usersTable          = "users"
	usersHaveRolesTable = "users_have_roles"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	params := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", params)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	return db, nil
}
