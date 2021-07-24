package user

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/parag08/go-microservice-boilerplate/pkg/logger"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

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

func (curp *CreateUserRequestParams) ValidateRequest() error {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		logger.Log.Fatal("translator not found")
	}

	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		logger.Log.Fatal(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
		return ut.Add("passwd", "{0} is not strong enough", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwd", fe.Field())
		return t
	})

	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	err := v.Struct(curp)

	return err
}
