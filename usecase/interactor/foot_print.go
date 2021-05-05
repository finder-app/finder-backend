package interactor

import (
	"finder/domain"
	"finder/interface/repository"
)

type FootPrintInteractor interface {
	GetFootPrintUsersByUid(currentUserUid string) ([]domain.FootPrint, error)
}

type footPrintInteractor struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintInteractor(ur repository.FootPrintRepository) *footPrintInteractor {
	return &footPrintInteractor{
		footPrintRepository: ur,
	}
}

func (i *footPrintInteractor) GetFootPrintUsersByUid(currentUserUid string) ([]domain.FootPrint, error) {
	footPrints, err := i.footPrintRepository.GetFootPrintUsersByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	// 取得したfoot_printsを使って、未読の足跡を全て既読にする！
	// if err := i.footPrintRepository.UpdateToAlreadyRead(currentUserUid); err != nil {
	// 	return nil, err
	// }
	return footPrints, nil
}
