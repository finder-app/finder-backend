package usecase

import (
	"context"
	"grpc/pb"
	"grpc/repository"
)

type FootPrintUsecase interface {
	// GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error)
	GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountReq) (*pb.GetUnreadCountRes, error)
}

type footPrintUsecase struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintUsecase(footPrintRepository repository.FootPrintRepository) FootPrintUsecase {
	return &footPrintUsecase{
		footPrintRepository: footPrintRepository,
	}
}

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

func (u *footPrintUsecase) GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountReq) (*pb.GetUnreadCountRes, error) {
	unreadCount, err := u.footPrintRepository.GetUnreadCount(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUnreadCountRes{
		UnreadCount: int64(unreadCount),
	}, nil
}
