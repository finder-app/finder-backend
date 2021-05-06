package interactor

import (
	"finder/domain"
	"finder/interface/repository"
	"fmt"
)

type FootPrintInteractor interface {
	GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error)
}

type footPrintInteractor struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintInteractor(ur repository.FootPrintRepository) *footPrintInteractor {
	return &footPrintInteractor{
		footPrintRepository: ur,
	}
}

func (i *footPrintInteractor) GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error) {
	if err := i.footPrintRepository.UpdateToAlreadyRead(currentUserUid); err != nil {
		fmt.Printf("interactor error %v", err)
		return nil, err
	}
	footPrints, err := i.footPrintRepository.GetFootPrintsByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return footPrints, nil
}