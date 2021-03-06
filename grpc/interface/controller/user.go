package controller

import (
	"context"

	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"
	"github.com/finder-app/finder-backend/grpc/interface/converter"
	"github.com/finder-app/finder-backend/grpc/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error) {
	currentUserUid := req.CurrentUserUid
	users, err := c.userUsecase.GetUsers(currentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersRes{
		Users: converter.ConvertUsers(users),
	}, nil
}

func (c *UserController) GetUserByUid(ctx context.Context, req *pb.GetUserByUidReq) (*pb.GetUserByUidRes, error) {
	uid := req.Uid
	// HACK: visitorIdよりcurrent_user_uidの方が良いのでは？ここは！でもfoot_print作るからなあ〜
	visitorUid := req.VisitorUid
	user, err := c.userUsecase.GetUserByUid(uid, visitorUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByUidRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *UserController) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	inputUser := &domain.User{
		Uid:       req.Uid,
		Email:     req.Email,
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Gender:    req.Gender,
	}
	user, err := c.userUsecase.CreateUser(inputUser)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserRes{
		User: converter.ConvertUser(user),
	}, nil
}
