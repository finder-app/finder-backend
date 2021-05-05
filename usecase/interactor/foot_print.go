package interactor

import (
	"finder/domain"
	"finder/interface/repository"
)

type FootPrintInteractor interface {
	GetFootPrintUsersByUid(uid string) ([]domain.User, error)
}

type footPrintInteractor struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintInteractor(ur repository.FootPrintRepository) *footPrintInteractor {
	return &footPrintInteractor{
		footPrintRepository: ur,
	}
}

func (i *footPrintInteractor) GetFootPrintUsersByUid(uid string) ([]domain.User, error) {
	users, err := i.footPrintRepository.GetFootPrintUsersByUid(uid)
	if err != nil {
		return nil, err
	}
	// 未読の足跡を全て既読にする！
	return users, nil
}
