package entity

import (
	"github.com/arilsonsantos/crud-go.git/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}

func UserDomainToEntity(domain model.UserDomainInterface) *UserEntity {
	return &UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func UserEntityToDomain(entity UserEntity) model.UserDomainInterface {
	userDomain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age)

	userDomain.SetID(entity.ID.Hex())
	return userDomain
}
