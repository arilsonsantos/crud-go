package entity

import (
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}

func UserDomainToEntity(domain domain.UserDomainInterface) *UserEntity {
	return &UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func UserEntityToDomain(entity UserEntity) domain.UserDomainInterface {
	userDomain := domain.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age)

	userDomain.SetID(entity.ID.Hex())
	return userDomain
}
