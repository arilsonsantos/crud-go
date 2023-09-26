package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_Delete(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := NewUserDomainService(repository)

	t.Run("delete_with_success", func(t *testing.T) {
		repository.EXPECT().Delete("123").Return(nil)
		err := service.Delete("123")

		assert.Nil(t, err)
	})

	t.Run("delete_with_error", func(t *testing.T) {
		repository.EXPECT().Delete("123").Return(errors.InternalServerError("Error trying to delete user"))
		err := service.Delete("123")

		assert.NotNil(t, err)

	})
}
