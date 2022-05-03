package repository

import (
	"errors"
	"fmt"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
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
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return 0, err
	}

	var id int

	CreateUserQuery := fmt.Sprintf("INSERT INTO %s (login, username, password_hash, created_at,profile_image) values ($1, $2, $3, $4,$5) RETURNING id", usersTable)
	row := tx.QueryRow(CreateUserQuery, user.Login,user.Username, user.Password,time.Now(),defaultAvatarUrl)
	err = row.Scan(&id)

	if err != nil {
		if strings.Contains(err.Error(),"duplicate key value"){
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return 0, errors.New("user already exist")
		}

		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return 0, err
	}

	CreateUserRoleQuery := fmt.Sprintf("INSERT INTO %s (user_id, role_id) values ($1, 2)", usersHaveRolesTable)
	_, err = tx.Exec(CreateUserRoleQuery, id)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id,login,password_hash,created_at,profile_image FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return user, err
	}

	query = fmt.Sprintf("SELECT role_id FROM %s WHERE user_id=$1", usersHaveRolesTable)
	err = r.db.Get(&user, query, user.Id)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return user, err
	}

	return user, nil
}
