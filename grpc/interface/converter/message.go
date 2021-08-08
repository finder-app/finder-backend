package converter

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertMessages(messages []*domain.Message) []*pb.Message {
	var pbMessages []*pb.Message
	for _, message := range messages {
		pbMessage := ConvertMessage(message)
		pbMessages = append(pbMessages, pbMessage)
	}
	return pbMessages
}

func ConvertMessage(message *domain.Message) *pb.Message {
	pbMessage := &pb.Message{
		Id:        message.Id,
		RoomId:    message.RoomId,
		UserUid:   message.UserUid,
		Text:      message.Text,
		Unread:    message.Unread,
		CreatedAt: timestamppb.New(message.CreatedAt),
		UpdatedAt: timestamppb.New(message.UpdatedAt),
	}
	return pbMessage
}
