package converter

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertLike(like *domain.Like) *pb.Like {
	pbLike := &pb.Like{
		Id:              like.Id,
		SentUserUid:     like.SentUserUid,
		SentUser:        ConvertUser(&like.SentUser),
		RecievedUserUid: like.RecievedUserUid,
		Skipped:         like.Skipped,
		Consented:       like.Consented,
		CreatedAt:       timestamppb.New(like.CreatedAt),
		UpdatedAt:       timestamppb.New(like.UpdatedAt),
	}
	return pbLike
}
