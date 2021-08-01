package controller

import (
	"context"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
)

type MessageController struct {
	// 	messageUsecase usecase.MessageUsecase
}

// func NewMessageController(messageUsecase usecase.MessageUsecase) *MessageController {
func NewMessageController() *MessageController {
	return &MessageController{
		// messageUsecase: messageUsecase,
	}
}

func (c *MessageController) GetMessages(ctx context.Context, req *pb.GetMessagesReq) (*pb.GetMessagesRes, error) {
	// messages, err := c.messageUsecase.GetMessages(req.CurrentUserUid)
	// if err != nil {
	// 	return nil, err
	// }
	messages := []*domain.Message{}
	message := &domain.Message{
		Id:      4545,
		RoomId:  100,
		UserUid: "hoge",
		Text:    "texttttttt",
		Unread:  false,
	}
	messages = append(messages, message)
	return &pb.GetMessagesRes{
		Messages: converter.ConvertMessages(messages),
	}, nil
}
