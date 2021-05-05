package interactor

import (
	"finder/domain"
	"finder/interface/repository"
)

type FootPrintInteractor interface {
	GetFootPrintsByUid(uid string) ([]domain.FootPrint, error)
}

type footPrintInteractor struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintInteractor(ur repository.FootPrintRepository) *footPrintInteractor {
	return &footPrintInteractor{
		footPrintRepository: ur,
	}
}

func (i *footPrintInteractor) GetFootPrintsByUid(uid string) ([]domain.FootPrint, error) {
	footPrints, err := i.footPrintRepository.GetFootPrintsByUid(uid)
	if err != nil {
		return nil, err

	}
	// 取得したfoot_printsを使って、未読の足跡を全て既読にする！
	// if err := i.footPrintRepository.UpdateFootPrint(&footPrints); err != nil {
	// 	return nil, err
	// }
	return footPrints, nil
}
