package handler

import (
	"context"
	"github.com/amrchnk/auth_service/pkg/models"
	pb "github.com/amrchnk/auth_service/proto"
)

func (i *Implementation) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	user, err := i.Service.GetUserById(req.Slug)
	if err != nil {
		return nil, err
	}
	userResp := pb.User{
		Slug:       user.Id,
		Login:      user.Login,
		Password:   user.Password,
		Username:   user.Username,
		UserRoleId: user.RoleId,
		ProfileImage: user.ProfileImage,
	}
	return &pb.GetUserByIdResponse{
		User: &userResp,
	}, err
}

func (i *Implementation) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	resp, err := i.Service.DeleteUserById(req.Slug)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserByIdResponse{
		Resp: resp,
	}, err
}

func (i *Implementation) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	userReq:=models.User{
		Login: req.User.Login,
		Id: req.User.Slug,
		Password: req.User.Password,
		Username: req.User.Username,
		ProfileImage: req.User.ProfileImage,
		RoleId: req.User.UserRoleId,
	}
	resp,err := i.Service.UpdateUser(userReq)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Resp: resp,
	}, err
}

func (i *Implementation) GetAllUsers(ctx context.Context,request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := i.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	usersResp := make([]*pb.User, 0, len(users))
	for i := range users {
		usersResp = append(usersResp, &pb.User{
			Slug:       users[i].Id,
			Login:      users[i].Login,
			Password:   users[i].Password,
			Username:   users[i].Username,
			UserRoleId: users[i].RoleId,
			ProfileImage: users[i].ProfileImage,
		})
	}

	return &pb.GetAllUsersResponse{
		User: usersResp,
	}, err
}

