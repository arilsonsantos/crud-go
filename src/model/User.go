package model

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
)

type User struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserInterface interface {
	Create(user User) *errors.ErrorDto
	Find(string) (*User, *errors.ErrorDto)
	Update(string, User) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
}
