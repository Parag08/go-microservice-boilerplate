package user

type CreateUserRequestParams struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,passwd"`
}

type User struct {
	UserID    string `json:"userId" bson:"userId"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}
