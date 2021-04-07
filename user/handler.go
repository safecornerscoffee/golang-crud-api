package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	db         *sql.DB
	repository UserRepository
}

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{
		db:         db,
		repository: UserRepository{db: db},
	}
}

func (u *userHandler) CreateUser(c echo.Context) (err error) {
	user := &User{}
	err = c.Bind(&user)
	if err != nil {
		return err
	}
	user, err = u.repository.CreateUser(*user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) GetUser(c echo.Context) (err error) {
	id := c.Param("id")
	user, err := u.repository.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found.")
	}
	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) FetchUsers(c echo.Context) (err error) {
	return c.String(http.StatusOK, "fetchUsers")
}

func (u *userHandler) UpdateUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "updateUser")
}

func (u *userHandler) DeleteUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "deleteUser")
}

func RegisterRoutes(db *sql.DB, e *echo.Echo) {
	h := NewUserHandler(db)

	e.GET("/users", h.FetchUsers)
	e.GET("/users/:id", h.GetUser)
	e.POST("/users", h.CreateUser)
	e.DELETE("/users/:id", h.DeleteUser)
	e.PUT("/users/:id", h.UpdateUser)
}
