package user

import "github.com/parag08/go-microservice-boilerplate/pkg/user"

type UserSerializer interface {
	Decode(input []byte) (*user.User, error)
	Encode(input *user.User) ([]byte, error)
}
