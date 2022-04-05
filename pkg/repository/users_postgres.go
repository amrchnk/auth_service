package repository

import (
	"fmt"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"sync"
)

var mu = &sync.Mutex{}

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserById(id int64) (models.User, error) {
	var user models.User

	mu.Lock()
	defer mu.Unlock()

	query := fmt.Sprintf("SELECT id, login, password_hash, username, created_at FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		log.Println("ERROR: ", err)
		return user, err
	}

	query = fmt.Sprintf("SELECT role_id FROM %s WHERE user_id=$1", usersHaveRolesTable)
	err = r.db.Get(&user, query, user.Id)
	if err != nil {
		log.Println("ERROR: ", err)
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) DeleteUserById(id int64) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("ERROR: ", err)
		return "ERROR: ", err
	}

	return fmt.Sprintf("User with id = %d was deleted successfully", id), nil
}

func (r *UserPostgres) UpdateUser(user models.User) (string, error) {

	mu.Lock()
	defer mu.Unlock()

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Login != "" {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, user.Login)
		argId++
	}

	if user.Password != "" {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, user.Password)
		argId++
	}

	if user.Username != "" {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, user.Username)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ut SET %s WHERE ut.id = %d`,
		usersTable, setQuery,user.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("ERROR: ", err)
		return "ERROR: ", err
	}

	if user.RoleId != 0 {
		query = fmt.Sprintf(`UPDATE %s urt SET role_id=$1 WHERE urt.user_id=$2`,
			usersHaveRolesTable)

		_, err = r.db.Exec(query, user.RoleId, user.Id)
		if err != nil {
			log.Println("ERROR: ", err)
			return "ERROR: ", err
		}
	}

	return fmt.Sprintf("user with id = %d was updated", user.Id), err
}

func (r *UserPostgres) GetAllUsers() ([]models.User, error) {
	mu.Lock()
	defer mu.Unlock()

	var users []models.User

	query := fmt.Sprintf("SELECT id, login, username, password_hash, created_at, role_id FROM %s u LEFT JOIN %s uhr on u.id=uhr.user_id", usersTable, usersHaveRolesTable)
	err := r.db.Select(&users, query)
	if err != nil {
		log.Println("ERROR: ", err)
		return users, err
	}

	return users, nil
}
