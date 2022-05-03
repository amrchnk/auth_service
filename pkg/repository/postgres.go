package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	usersTable          = "users"
	usersHaveRolesTable = "users_have_roles"

	defaultAvatarUrl = "https://res.cloudinary.com/disfhw1xf/image/upload/v1651590413/design_app/avatars/defaut_profile_image_xqvcqa.png"
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

	err = db.Ping()
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	return db, nil
}
