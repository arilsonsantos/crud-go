package view

import (
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConvertUserDomainToUserDto(t *testing.T) {
	userDomain := domain.NewUserDomain(
		"test@email.com",
		"!23456",
		"John Test",
		42,
	)

	userResponseDto := ConvertUserDomainToUserDto(userDomain)

	assert.EqualValues(t, reflect.TypeOf(userResponseDto),
		reflect.TypeOf(dto.UserResponseDto{}))
	assert.EqualValues(t, userResponseDto.Email, userDomain.GetEmail())
}
