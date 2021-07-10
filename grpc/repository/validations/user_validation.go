package validations

import (
	"errors"
	"grpc/domain"

	"github.com/go-playground/validator"
)

// NOTE: 今は使用していない
func ValidateUser(user *domain.User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	if err := validateUserGender(user); err != nil {
		return err
	}
	return nil
}

func validateUserGender(user *domain.User) error {
	// NOTE: 2パターン考えたけど、前者のほうが分かりやすい
	if user.Gender == "男性" || user.Gender == "女性" {
		return nil
	}
	return errors.New("genderが不正な値です")

	// if user.Gender != "男性" && user.Gender != "女性" {
	// 	return errors.New("genderが不正な値です")
	// }
	// return nil
}
