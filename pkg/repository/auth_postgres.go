package repository

import (
	"fmt"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println("ERROR: ", err)
		return 0, err
	}

	var id int

	CreateUserQuery := fmt.Sprintf("INSERT INTO %s (login, password_hash,created_at) values ($1, $2, $3) RETURNING id", usersTable)
	row := tx.QueryRow(CreateUserQuery, user.Login, user.Password,time.Now())
	err = row.Scan(&id)
	if err != nil {
		log.Println("ERROR: ", err)
		return 0, err
	}

	CreateUserRoleQuery := fmt.Sprintf("INSERT INTO %s (user_id, role_id) values ($1, 2)", usersHaveRolesTable)
	_, err = tx.Exec(CreateUserRoleQuery, id)
	if err != nil {
		log.Println("ERROR: ", err)
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id,login,password_hash FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
