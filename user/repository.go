package user

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

const (
	insertUserSQL = `Insert Into users (id, name, email, password, role)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	getUserSQL    = `SELECT id, name, email FROM users WHERE id=$1`
	updateUserSQL = `UPDATE users SET name=$2, email=$3 WHERE id=$1`
	deleteUserSQL = `DELETE FROM users where id=$1`
)

func (u *UserRepository) updateUser(user User) (err error) {
	_, err = u.db.Query(updateUserSQL, user.ID, user.Name, user.Email)
	return err
}

func (u *UserRepository) CreateUser(user User) (*User, error) {
	user.ID = uuid.New().String()

	var err error
	user.Password, err = hashAndSaltPwd(user.Password)
	if err != nil {
		return nil, err
	}

	err = u.db.QueryRow(insertUserSQL, user.ID, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (u *UserRepository) Get(id string) (User, error) {
	user := User{}
	row := u.db.QueryRow(getUserSQL, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (u *UserRepository) Update(user User) error {
	return u.updateUser(user)
}

func (u *UserRepository) Delete(id string) error {
	_, err := u.db.Query(deleteUserSQL, id)
	return err
}

func hashAndSaltPwd(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}
