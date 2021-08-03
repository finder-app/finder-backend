package controller

import (
	"context"
	"fmt"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/usecase"
	"io"
	"os"
)

type ProfileController struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileController(profileUsecase usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: profileUsecase,
	}
}

func (c *ProfileController) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error) {
	user, err := c.profileUsecase.GetProfile(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *ProfileController) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileRes, error) {
	// NOTE: 変更して良いカラムだけ書く uidは更新時にwhereするのに必要
	inputUser := &domain.User{
		Uid:       req.User.Uid,
		LastName:  req.User.LastName,
		FirstName: req.User.FirstName,
	}
	user, err := c.profileUsecase.UpdateProfile(inputUser)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *ProfileController) TestImage(stream pb.ProfileService_TestImageServer) error {
	file, err := os.Create("tmp2.png")
	defer file.Close()
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		file.Write(resp.Image)
	}
	fmt.Println("filefilefilefile", file)
	err = stream.SendAndClose(&pb.TestImageRes{Response: "OK"})
	if err != nil {

		return err
	}
	return nil
	// var pointCount, featureCount, distance int32
	// var lastPoint *pb.Point
	// startTime := time.Now()
	// for {
	// 	point, err := stream.Recv()
	// 	if err == io.EOF {
	// 		endTime := time.Now()
	// 		return stream.SendAndClose(&pb.RouteSummary{
	// 			PointCount:   pointCount,
	// 			FeatureCount: featureCount,
	// 			Distance:     distance,
	// 			ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
	// 		})
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// 	pointCount++
	// 	for _, feature := range s.savedFeatures {
	// 		if proto.Equal(feature.Location, point) {
	// 			featureCount++
	// 		}
	// 	}
	// 	if lastPoint != nil {
	// 		distance += calcDistance(lastPoint, point)
	// 	}
	// 	lastPoint = point
	// }
}
