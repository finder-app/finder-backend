package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
)

type FootPrintUsecase interface {
	GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error)
}

type footPrintUsecase struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintUsecase(ur repository.FootPrintRepository) *footPrintUsecase {
	return &footPrintUsecase{
		footPrintRepository: ur,
	}
}

func (i *footPrintUsecase) GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error) {
	if err := i.footPrintRepository.UpdateToAlreadyRead(currentUserUid); err != nil {
		return nil, err
	}
	footPrints, err := i.footPrintRepository.GetFootPrintsByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return footPrints, nil
}
