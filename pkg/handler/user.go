package handler

import (
	"context"
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
		})
	}

	return &pb.GetAllUsersResponse{
		User: usersResp,
	}, err
}
