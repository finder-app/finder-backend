package controller

import (
	"context"

	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"
	"github.com/finder-app/finder-backend/grpc/interface/converter"
	"github.com/finder-app/finder-backend/grpc/usecase"
)

type FootPrintController struct {
	footPrintUsecase usecase.FootPrintUsecase
}

func NewFootPrintController(footPrintUsecase usecase.FootPrintUsecase) *FootPrintController {
	return &FootPrintController{
		footPrintUsecase: footPrintUsecase,
	}
}

func (c *FootPrintController) GetFootPrints(ctx context.Context, req *pb.GetFootPrintsReq) (*pb.GetFootPrintsRes, error) {
	footPrints, err := c.footPrintUsecase.GetFootPrints(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetFootPrintsRes{
		FootPrints: converter.ConvertFootPrints(footPrints),
	}, nil
}

func (c *FootPrintController) GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountReq) (*pb.GetUnreadCountRes, error) {
	unreadCount, err := c.footPrintUsecase.GetUnreadCount(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUnreadCountRes{
		UnreadCount: unreadCount,
	}, nil
}
