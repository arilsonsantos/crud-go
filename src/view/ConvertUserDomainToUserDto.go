package view

import (
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/model"
)

func ConvertUserDomainToUserDto(
	userDomain model.UserDomainInterface,
) dto.UserResponseDto {
	return dto.UserResponseDto{
		Id:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
