package handler

import (
	"github.com/amrchnk/auth_service/pkg/service"
	pb "github.com/amrchnk/auth_service/proto"
	"sync"
)

type Implementation struct {
	pb.UnimplementedAuthServiceServer
	*service.Service
	mu sync.Mutex
}

func NewAuthService(s *service.Service)*Implementation{
	return &Implementation{
		Service:s,
	}
}
