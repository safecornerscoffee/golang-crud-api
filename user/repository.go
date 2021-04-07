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
	insertUserSQL = `Insert Into users (id, name, email, password, role
		VALUES ($1, $2, $3, $4, $5) returning id`
	getUserSQL    = `SELECT id, name, email, role FROM users WHERE id = $1`
	updateUserSQL = `UPDATE users SET %s WHERE id = $1`
	deleteUserSQL = `DELETE FROM users where id = $1`
)

func (u *UserRepository) updateUser(user User, fields ...string) (err error) {

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
	return &user, err
}

func (u *UserRepository) Get(id string) (User, error) {
	user := User{}
	var err error
	return user, err
}

func (u *UserRepository) Update(user User) error {
	var err error
	return err
}

func (u *UserRepository) Delete(id string) error {
	var err error
	return err
}

func hashAndSaltPwd(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}
