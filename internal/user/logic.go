package user

import (
	"errors"

	"github.com/parag08/go-microservice-boilerplate/internal/common"
	"github.com/parag08/go-microservice-boilerplate/pkg/user"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (u *userService) Find(ctx common.Context, userEmail string) *user.User {
	return u.userRepo.Find(ctx, userEmail)
}

func (u *userService) Store(ctx common.Context, params *user.CreateUserRequestParams) error {
	return u.userRepo.Store(ctx, params)
}
