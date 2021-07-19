package validation

import (
	"errors"
	"grpc/domain"
)

func ValidateLike(like *domain.Like) error {
	// FIXME: likeのfieldにSentUserがあり、それがvalidatorに引っかかるためコメントアウト
	// validate := validator.New()
	// if err := validate.Struct(like); err != nil {
	// 	return err
	// }
	if err := consentedOrSkipped(like); err != nil {
		return err
	}
	return nil
}

// NOTE: いいね済みかskip済みならerrorを返す
func consentedOrSkipped(like *domain.Like) error {
	if like.Consented {
		return errors.New("consnted like")
	} else if like.Skipped {
		return errors.New("skipped like")
	}
	return nil
}
