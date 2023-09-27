package repository

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

var testEmail = "test@email.com"
var testName = "John Test"
var testAge int8 = 42
var errorDBMessage = "Error from database"

func TestUserRepositoryInterface_Create(t *testing.T) {
	mtestDb, databaseName, _ := getDataBaseTest(t)

	mtestDb.Run("create success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain := domain.NewUserDomain(testEmail, "123", testName, testAge)
		userDomain, err := repo.Create(userDomain)
		id, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.NotNil(t, id)
		assert.EqualValues(t, userDomain.GetEmail(), testEmail)
		assert.EqualValues(t, userDomain.GetPassword(), "123")
		assert.EqualValues(t, userDomain.GetName(), testName)
		assert.EqualValues(t, userDomain.GetAge(), testAge)
	})

	mtestDb.Run(errorDBMessage, func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain := domain.NewUserDomain(testEmail, "123", testName, testAge)
		userDomainReturn, _ := repo.Create(userDomain)

		assert.Nil(t, userDomainReturn)
	})

}

func TestUserRepositoryInterface_FindByEmail(t *testing.T) {
	mTestDb, databaseName, collectionName := getDataBaseTest(t)

	mTestDb.Run("when_sending_a_valid_email_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    testName,
			Password: "123",
			Name:     "John Lennon",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)

	})

	mTestDb.Run("when_error_in_mongodb_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmail("teste@email.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "Error trying to find user by email")
	})

	mTestDb.Run("when_no_content_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			// no documents listed here on purpose
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmail("teste@email.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "User not found with this email: teste@email.com")

	})
}

func TestUserRepositoryInterface_FindById(t *testing.T) {
	mTestDb, databaseName, collectionName := getDataBaseTest(t)

	mTestDb.Run("when_sending_a_valid_id_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    testName,
			Password: "123",
			Name:     "John Lennon",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindById(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mTestDb.Run("when_error_in_mongodb_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindById(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "Error trying to find user by ID")
	})

	mTestDb.Run("when_no_content_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			// no documents listed here on purpose
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindById("123")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "User not found with this ID: 123")
	})
}

func TestUserRepositoryInterface_FindByEmailAndPassword(t *testing.T) {
	mTestDb, databaseName, collectionName := getDataBaseTest(t)

	mTestDb.Run("when_sending_a_valid_email_and_password_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Email:    testName,
			Password: "123",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mTestDb.Run("when_error_in_mongodb_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmailAndPassword("email", "password")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "Error trying to find user login by email and password")
	})

	mTestDb.Run("when_no_content_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			// no documents listed here on purpose
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain, err := repo.FindByEmailAndPassword("email", "password")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.EqualError(t, err, "Email or password is invalid.")
	})
}

func TestUserRepositoryInterface_Delete(t *testing.T) {
	mTestDb, databaseName, _ := getDataBaseTest(t)

	mTestDb.Run("delete success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		err := repo.Delete("123")

		assert.Nil(t, err)
	})

	mTestDb.Run(errorDBMessage, func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		err := repo.Delete("123")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Error trying to delete user")
	})

}

func TestUserRepositoryInterface_Update(t *testing.T) {
	mTestDb, databaseName, _ := getDataBaseTest(t)

	mTestDb.Run("update success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain := domain.NewUserUpdateDomain(testName, testAge)
		userDomain.SetID(primitive.NewObjectID().Hex())
		err := repo.Update(userDomain.GetID(), userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetName(), testName)
		assert.EqualValues(t, userDomain.GetAge(), testAge)
	})

	mTestDb.Run(errorDBMessage, func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepositoryInterface(databaseMock)
		userDomain := domain.NewUserUpdateDomain(testName, testAge)
		userDomain.SetID(primitive.NewObjectID().Hex())
		err := repo.Update(userDomain.GetID(), userDomain)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Error trying to update user")
	})

}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "name", Value: userEntity.Name},
		{Key: "password", Value: userEntity.Password},
		{Key: "age", Value: userEntity.Age},
	}
}

func getDataBaseTest(t *testing.T) (*mtest.T, string, string) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()
	return mTestDb, databaseName, collectionName
}
