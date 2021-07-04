package usecase

// import (
// 	"finder/domain"
// 	"finder/infrastructure/repository"
// )

// type FootPrintUsecase interface {
// 	GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error)
// 	GetUnreadCount(currentUserUid string) (int, error)
// }

// type footPrintUsecase struct {
// 	footPrintRepository repository.FootPrintRepository
// }

// func NewFootPrintUsecase(ur repository.FootPrintRepository) *footPrintUsecase {
// 	return &footPrintUsecase{
// 		footPrintRepository: ur,
// 	}
// }

// func (u *footPrintUsecase) GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error) {
// 	if err := u.footPrintRepository.UpdateToAlreadyRead(currentUserUid); err != nil {
// 		return nil, err
// 	}
// 	footPrints, err := u.footPrintRepository.GetFootPrintsByUid(currentUserUid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return footPrints, nil
// }

// func (u *footPrintUsecase) GetUnreadCount(currentUserUid string) (int, error) {
// 	return u.footPrintRepository.GetUnreadCount(currentUserUid)
// }
