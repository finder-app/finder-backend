package usecase

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/repository"
)

type FootPrintUsecase interface {
	GetFootPrints(currentUserUid string) ([]*domain.FootPrint, error)
	GetUnreadCount(currentUserUid string) (int64, error)
}

type footPrintUsecase struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintUsecase(footPrintRepository repository.FootPrintRepository) FootPrintUsecase {
	return &footPrintUsecase{
		footPrintRepository: footPrintRepository,
	}
}

func (u *footPrintUsecase) GetFootPrints(currentUserUid string) ([]*domain.FootPrint, error) {
	if err := u.footPrintRepository.UpdateToAlreadyRead(currentUserUid); err != nil {
		return nil, err
	}
	return u.footPrintRepository.GetFootPrintsByUid(currentUserUid)
}

func (u *footPrintUsecase) GetUnreadCount(currentUserUid string) (int64, error) {
	return u.footPrintRepository.GetUnreadCount(currentUserUid)
}
