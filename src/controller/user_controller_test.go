package controller

import (
	"encoding/json"
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var testEmail = "test@email.com"
var errorMessageTest = "Error test"
var testName = "Jonh Test"
var testAge int8 = 42

// Find
func TestUserControllerInterface_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("email_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userEmail", Value: "test_error"},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userEmail", Value: testEmail},
		}

		mockService.EXPECT().FindByEmail(testEmail).Return(
			nil, errors.InternalServerError(errorMessageTest))
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userEmail", Value: testEmail},
		}

		userDomain := domain.NewUserDomain(testEmail, "123", "Test", testAge)
		mockService.EXPECT().FindByEmail(testEmail).Return(
			userDomain, nil)
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestUserControllerInterface_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := "123"

		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindById(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		mockService.EXPECT().FindById(id).Return(
			nil, errors.InternalServerError(errorMessageTest))
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindById(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		userDomain := domain.NewUserDomain(testEmail, "123", "Test", testAge)
		mockService.EXPECT().FindById(id).Return(
			userDomain, nil)
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindById(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestUserControllerInterface_FindUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserLogin{
			Email:    "invalid_email",
			Password: "123",
		}

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.FindUserLogin(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("it_is_valid_but_server_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserLogin{
			Email:    "valid@email.com",
			Password: "!23456",
		}

		userDomain := domain.NewUserLoginDomain(userRequest.Email, userRequest.Password)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().LoginUserService(userDomain).Return(nil, "", errors.InternalServerError(errorMessageTest))

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.FindUserLogin(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("it_is_valid_server_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserLogin{
			Email:    "valid@email.com",
			Password: "!23456",
		}

		userDomain := domain.NewUserLoginDomain(userRequest.Email, userRequest.Password)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().LoginUserService(userDomain).Return(userDomain, "", nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.FindUserLogin(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

// Delete
func TestUserControllerInterface_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := "123"
		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.Delete(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		mockService.EXPECT().Delete(id).Return(errors.InternalServerError(errorMessageTest))
		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.Delete(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		mockService.EXPECT().Delete(id).Return(nil)
		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.Delete(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

// Create
func TestUserControllerInterface_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserRequestDto{
			Email:    "test",
			Name:     testName,
			Age:      0,
			Password: "123",
		}

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserRequestDto{
			Email:    testEmail,
			Name:     testName,
			Age:      testAge,
			Password: "123456!",
		}

		userDomain := domain.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().Create(userDomain).Return(nil, errors.InternalServerError(
			"Error trying to create user."))

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("it_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserRequestDto{
			Email:    testEmail,
			Name:     testName,
			Age:      testAge,
			Password: "123456!",
		}

		userDomain := domain.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().Create(userDomain).Return(userDomain, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPost, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})
}

// Update
func TestUserControllerInterface_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserDomainServiceInterface(ctrl)
	controller := NewUserControllerInterface(mockService)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserUpdateRequestDto{
			Name: "",
			Age:  -1,
		}

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		MakeRequest(context, []gin.Param{}, url.Values{}, http.MethodPut, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("userId_is_invalid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserUpdateRequestDto{
			Name: testName,
			Age:  testAge,
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		MakeRequest(context, param, url.Values{}, http.MethodPost, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("it_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserUpdateRequestDto{
			Name: testName,
			Age:  testAge,
		}

		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateDomain := domain.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().Update(id, updateDomain).Return(
			errors.InternalServerError("Error trying to call userService update"))

		MakeRequest(context, param, url.Values{}, http.MethodPut, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("it_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := dto.UserUpdateRequestDto{
			Name: testName,
			Age:  testAge,
		}

		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateDomain := domain.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		body, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		mockService.EXPECT().Update(id, updateDomain).Return(nil)

		MakeRequest(context, param, url.Values{}, http.MethodPut, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
