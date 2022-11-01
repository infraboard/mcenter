package validate

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func Validate(obj any) error {
	return validate.Struct(obj)
}
