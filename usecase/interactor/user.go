package interactor

import (
	"finder/domain"
	"finder/interface/controller"
	"finder/usecase/repository"
	"strconv"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) *userInteractor {
	return &userInteractor{ur}
}

func (i *userInteractor) GetUserByID(c controller.Context, userID int64) (*domain.User, error) {
	userID, _ := strconv.Atoi(c.Param("id"))
	result, err := i.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
