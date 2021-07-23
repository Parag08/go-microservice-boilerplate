package user

import (
	"github.com/parag08/go-microservice-boilerplate/internal/common"
	"github.com/parag08/go-microservice-boilerplate/pkg/user"
)

type UserService interface {
	Find(ctx common.Context, userEmail string) *user.User
	Store(ctx common.Context, user *user.CreateUserRequestParams) error
}
