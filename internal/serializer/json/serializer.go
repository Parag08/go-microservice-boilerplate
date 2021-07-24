package json

import (
	"encoding/json"

	"github.com/parag08/go-microservice-boilerplate/pkg/user"
	"github.com/pkg/errors"
)

type User struct{}

func (u *User) Decode(input []byte) (*user.User, error) {
	user := &user.User{}
	if err := json.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.User.Decode")
	}
	return user, nil
}

func (u *User) Encode(input *user.User) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}
