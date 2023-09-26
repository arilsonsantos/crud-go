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

func TestUserDomainService_Create(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain("test@email", "123", "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(userDomain, nil)
		user, err := service.Create(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account.")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain("test@email", "123", "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().Create(userDomain).Return(nil, errors.InternalServerError("Error trying to create user"))
		user, err := service.Create(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain("test@email", "123", "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().Create(userDomain).Return(userDomain, nil)
		user, err := service.Create(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})

}
