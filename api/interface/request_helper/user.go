package request_helper

import "api/pb"

// NOTE: 必須カラムはbinding: requiredしてる。Emailはなくてもエラー吐かない
type RequestUser struct {
	Uid       string `form:"uid" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	FirstName string `form:"first_name" binding:"required"`
	Email     string `form:"email"`
	Gender    string `form:"gender"`
}

func NewRequestUser() *RequestUser {
	return &RequestUser{}
}

func NewPbUser(requestUser *RequestUser) *pb.User {
	return &pb.User{
		Uid:       requestUser.Uid,
		LastName:  requestUser.LastName,
		FirstName: requestUser.FirstName,
		Email:     requestUser.Email,
		Gender:    requestUser.Gender,
	}
}
