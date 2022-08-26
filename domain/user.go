package domain

import (
	"time"

	"github.com/labstack/echo"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCase interface {
	Add(newUser User) (row int, err error)
	GetProfile(id int) (User, error)
}

type UserData interface {
	Insert(newUser User) (row int, err error)
	Get(id int) (User, error)
}

type UserHandler interface {
	InsertUser() echo.HandlerFunc
	GetUser() echo.HandlerFunc
}
