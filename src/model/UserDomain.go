package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/arilsonsantos/crud-go.git/src/errors"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (user *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	Create() *errors.ErrorDto
	Find(string) (*UserDomain, *errors.ErrorDto) // Update the return type here
	Update(string) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
}
