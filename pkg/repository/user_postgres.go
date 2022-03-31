package repository

/*import (
	"fmt"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserById(id int64) (models.User, error) {
	var user models.User

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
}*/

/*func (r *UserPostgres) DeleteUserById(id int64) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_,err := r.db.Exec(query, id)
	if err != nil {
		log.Println("ERROR: ", err)
		return "ERROR: ", err
	}

	return fmt.Sprintf("User with id = %d was deleted successfully",id), nil
}

func (r *UserPostgres) GetAllUsers() ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		log.Println("ERROR: ", err)
		return users, err
	}

	query = fmt.Sprintf("SELECT role_id FROM %s", usersHaveRolesTable)
	err = r.db.Select(&users, query)
	if err != nil {
		log.Println("ERROR: ", err)
		return users, err
	}

	return users, nil
}*/
