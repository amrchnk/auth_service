package handler

import (
	"context"
	"github.com/amrchnk/auth_service/pkg/models"
	pb "github.com/amrchnk/auth_service/proto"
	"github.com/spf13/cast"
)

func (i *Implementation) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error){
	var input models.User
	input.Login,input.Password=req.Login,req.Password
	id,err:=i.Service.CreateUser(input)
	if err!=nil{
		return nil, err
	}
	return &pb.SignUpResponse{
		Slug: cast.ToInt64(id),
	},nil
}

func (i *Implementation) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error){
	user,err:=i.Service.CheckUser(req.Login,req.Password)
	if err!=nil{
		return nil, err
	}
	return &pb.SignInResponse{
		Session: user,
	},nil
}
