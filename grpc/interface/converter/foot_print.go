package converter

import (
	"grpc/domain"
	"grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertFootPrints(footPrints []*domain.FootPrint) []*pb.FootPrint {
	var pbFootPrints []*pb.FootPrint
	for _, footPrint := range footPrints {
		pbFootPrint := ConvertFootPrint(footPrint)
		pbFootPrints = append(pbFootPrints, pbFootPrint)
	}
	return pbFootPrints
}

func ConvertFootPrint(footPrint *domain.FootPrint) *pb.FootPrint {
	pbFootPrint := &pb.FootPrint{
		Id:         footPrint.Id,
		VisitorUid: footPrint.VisitorUid,
		Visitor:    ConvertUser(footPrint.Visitor),
		Unread:     footPrint.Unread,
		CreatedAt:  timestamppb.New(footPrint.CreatedAt),
		UpdatedAt:  timestamppb.New(footPrint.UpdatedAt),
		UesrUid:    footPrint.UserUid,
	}
	return pbFootPrint
}
