package mocks

import (
	"api/finder-protocol-buffers/pb"
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UserClient struct {
	mock.Mock
}

func (u *UserClient) GetUsers(ctx context.Context, in *pb.GetUsersReq, opts ...grpc.CallOption) (*pb.GetUsersRes, error) {
	arguments := u.Called(ctx, in)

	var getUsersRes *pb.GetUsersRes
	if arguments.Get(0) != nil {
		fmt.Println(arguments)        // [*pb.GetUsersRes <nil>]
		fmt.Println(arguments.Get(0)) // *pb.GetUsersRes

		// FIXME: コメントインすると ↓のエラーを吐くためコメントアウト
		// interface conversion: interface {} is string, not *pb.GetUsersRes
		// getUsersRes = arguments.Get(0).(*pb.GetUsersRes)
	}
	err := MockArgumentsError(arguments, 1)
	return getUsersRes, err
}

func (u *UserClient) GetUserByUid(ctx context.Context, in *pb.GetUserByUidReq, opts ...grpc.CallOption) (*pb.GetUserByUidRes, error) {
	arguments := u.Called()

	// user := &domain.User{}
	var getUserByUidRes *pb.GetUserByUidRes
	// if arguments.Get(0) != nil {
	// 	user = arguments.Get(0).(*domain.User)
	// }
	err := MockArgumentsError(arguments, 1)
	return getUserByUidRes, err
}

func (u *UserClient) CreateUser(ctx context.Context, in *pb.CreateUserReq, opts ...grpc.CallOption) (*pb.CreateUserRes, error) {
	arguments := u.Called()
	var createUserRes *pb.CreateUserRes
	if arguments.Get(0) != nil {
		// user = arguments.Get(0).(*domain.User)
	}
	err := MockArgumentsError(arguments, 1)
	return createUserRes, err
}
