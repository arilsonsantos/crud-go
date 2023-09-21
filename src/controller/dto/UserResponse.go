package dto

import (
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
)

type UserResponseDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}

func ConvertUserDomainToUserDto(
	userDomain domain.UserDomainInterface,
) UserResponseDto {
	return UserResponseDto{
		Id:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
