package user

import (
	"github.com/parag08/go-microservice-boilerplate/internal/common"
	"github.com/parag08/go-microservice-boilerplate/pkg/user"
)

// Don't ask me why this is called repository, If it was upto me I would name this db.go

type UserRepository interface {
	Find(ctx common.Context, userEmail string) *user.User
	Store(ctx common.Context, user *user.CreateUserRequestParams) error
}
