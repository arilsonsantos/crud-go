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

var testName = "John Test"
var testPassword = "123"
var testAge int8 = 42
var errorMessage = "Error trying to find user by email and password."

func TestUserDomainService_LoginUserService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := &userDomainService{repository}

	t.Run("login_user_with_error_using_gomock", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain(testEmail, testPassword, testName, testAge)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(nil, errors.InternalServerError(
			errorMessage))
		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, errorMessage)
	})

	t.Run("login_user_with_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain(testEmail, testPassword, testName, testAge)
		userDomain.SetID(id)

		userDomainMock := domain.NewUserDomain(userDomain.GetEmail(), userDomain.GetPassword(), userDomain.GetName(), userDomain.GetAge())
		userDomainMock.EncryptPassword()

		repository.EXPECT().FindByEmailAndPassword(
			userDomain.GetEmail(), userDomainMock.GetPassword()).Return(nil, errors.InternalServerError(
			errorMessage))
		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, errorMessage)
	})

	t.Run("login_user_when_generate_token_error", func(t *testing.T) {
		mockUserDomainInterce := mocks.NewMockUserDomainInterface(control)
		mockUserDomainInterce.EXPECT().GetEmail().Return(testEmail)
		mockUserDomainInterce.EXPECT().GetPassword().Return(testPassword)
		mockUserDomainInterce.EXPECT().EncryptPassword()
		mockUserDomainInterce.EXPECT().GenerateToken().Return("", errors.InternalServerError(
			"Error trying to create token."))

		repository.EXPECT().FindByEmailAndPassword("test@email.com", testPassword).Return(
			mockUserDomainInterce, nil)
		user, token, err := service.LoginUserService(mockUserDomainInterce)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to create token.")
	})

	t.Run("login_user_when_generate_token_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := domain.NewUserDomain(testEmail, testPassword, testName, testAge)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmailAndPassword(userDomain.GetEmail(), gomock.Any()).Return(userDomain, nil)

		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotNil(t, token)
	})

}
