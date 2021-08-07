package controller

import (
	"context"
	"grpc/finder-protocol-buffers/pb"
	"grpc/interface/converter"
	"grpc/usecase"
)

type MessageController struct {
	messageUsecase usecase.MessageUsecase
}

func NewMessageController(messageUsecase usecase.MessageUsecase) *MessageController {
	return &MessageController{
		messageUsecase: messageUsecase,
	}
}

func (c *MessageController) GetMessages(ctx context.Context, req *pb.GetMessagesReq) (*pb.GetMessagesRes, error) {
	messages, err := c.messageUsecase.GetMessages(req.RoomId, req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetMessagesRes{
		Messages: converter.ConvertMessages(messages),
	}, nil
}
