package domain

import "github.com/arilsonsantos/crud-go.git/src/errors"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
	SetID(string2 string)
	GetID() string
	GenerateToken() (string, *errors.ErrorDto)
}
