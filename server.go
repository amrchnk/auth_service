package auth_service

import (
	"context"
	"encoding/json"
	services "github.com/amrchnk/auth_service/pkg/service"
	pb "github.com/amrchnk/auth_service/proto"
	"github.com/gin-contrib/sessions"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	pb.AuthServiceServer
	mu sync.Mutex
}

func (s *Server) SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	id, err := services.Authorization.CreateUser(proto.)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = id
	userSession, err := json.Marshal(input)
	if err != nil {
		return
	}
	session := sessions.Default(c)
	session.Set("UserSession",userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

/*func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}*/
