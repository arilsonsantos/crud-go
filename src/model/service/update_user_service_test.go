package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

var testEmail = "test@email.com"
var internalServerErrorMessage = "Error trying to update user"

func TestUserDomainService_Update(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := NewUserDomainService(repository)

	t.Run("when_send_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain(testEmail, "123", "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().Update(id, userDomain).Return(nil)
		err := service.Update(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_send_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain(testEmail, "123", "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().Update(id, userDomain).Return(errors.InternalServerError(internalServerErrorMessage))
		err := service.Update(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, internalServerErrorMessage)
	})
}
