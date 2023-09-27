package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestUserDomainService_FindById(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := NewUserDomainService(repository)

	t.Run("findById_success", func(t *testing.T) {
		repository.EXPECT().FindById("123").Return(domain.NewUserDomain(
			testEmail, testPassword, testName, testAge), nil)

		user, err := service.FindById(testPassword)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetEmail(), testEmail)
	})

	t.Run("findBydId_not_found", func(t *testing.T) {
		repository.EXPECT().FindById(testPassword).Return(nil, errors.NotFoundError("User not found with this ID: 123"))

		_, err := service.FindById(testPassword)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Code, http.StatusNotFound)
		assert.EqualError(t, err, "User not found with this ID: 123")
	})
}

func TestUserDomainService_FindEmail(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := NewUserDomainService(repository)

	t.Run("findByEmail_success", func(t *testing.T) {
		repository.EXPECT().FindByEmail(testEmail).Return(domain.NewUserDomain(
			testEmail, testPassword, "Teste", 42), nil)

		user, err := service.FindByEmail(testEmail)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetEmail(), testEmail)
	})

	t.Run("findByEmail_not_found", func(t *testing.T) {
		repository.EXPECT().FindByEmail(testEmail).Return(nil, errors.NotFoundError("User not found with this email: test@email"))

		_, err := service.FindByEmail(testEmail)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Code, http.StatusNotFound)
		assert.EqualError(t, err, "User not found with this email: test@email")
	})

}

func TestUserDomainService_FindByUserAndPassword(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepositoryInterface(control)
	service := &userDomainService{repository}

	t.Run("find_user_by_email_and_password_with_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := testEmail
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := domain.NewUserDomain(email, password, "Test Name", 42)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findByEmailAndPassword(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), userDomain.GetID())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())

	})

	t.Run("find_user_by_email_and_password_not_found", func(t *testing.T) {
		email := testEmail
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindByEmailAndPassword(email, password).Return(nil, errors.NotFoundError("User not found."))

		_, err := service.findByEmailAndPassword(email, password)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Code, http.StatusNotFound)
		assert.EqualValues(t, err.Message, "User not found.")

	})

}
