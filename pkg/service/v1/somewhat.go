package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/dhanusaputra/somewhat-server/util/authutil"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// Server ...
type Server struct {
	data     map[string]interface{}
	userData []v1.User
}

// NewServer ...
func NewServer(data map[string]interface{}, userData []v1.User) *Server {
	return &Server{
		data:     data,
		userData: userData,
	}
}

// GetSomething ...
func (s *Server) GetSomething(ctx context.Context, req *v1.GetSomethingRequest) (*v1.GetSomethingResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Id]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to find ID: %v", req.Id))
	}
	b, err := json.Marshal(s.data[req.Id])
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("json has invalid format, err: %v", err))
	}
	return &v1.GetSomethingResponse{
		Api: apiVersion,
		Something: &v1.Something{
			Id:          req.Id,
			Description: string(b),
		},
	}, nil
}

// CreateSomething ...
func (s *Server) CreateSomething(ctx context.Context, req *v1.CreateSomethingRequest) (*v1.CreateSomethingResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if strings.TrimSpace(req.Something.Id) == "" {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID is required"))
	}
	if _, ok := s.data[req.Something.Id]; ok {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("failed to create, ID: %v already exists", req.Something.Id))
	}
	var desc map[string]string
	err := json.Unmarshal([]byte(req.Something.Description), &desc)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("json has invalid format, err: %v", err))
	}
	s.data[req.Something.Id] = desc
	return &v1.CreateSomethingResponse{
		Api: apiVersion,
		Id:  req.Something.Id,
	}, nil
}

// UpdateSomething ...
func (s *Server) UpdateSomething(ctx context.Context, req *v1.UpdateSomethingRequest) (*v1.UpdateSomethingResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Something.Id]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to find ID: %v", req.Something.Id))
	}
	var desc map[string]string
	err := json.Unmarshal([]byte(req.Something.Description), &desc)
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("json has invalid format, err: %v", err))
	}
	s.data[req.Something.Id] = desc
	return &v1.UpdateSomethingResponse{
		Api:     apiVersion,
		Updated: true,
	}, nil
}

// DeleteSomething ...
func (s *Server) DeleteSomething(ctx context.Context, req *v1.DeleteSomethingRequest) (*v1.DeleteSomethingResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Id]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to find ID: %v", req.Id))
	}
	delete(s.data, req.Id)
	return &v1.DeleteSomethingResponse{
		Api:     apiVersion,
		Deleted: true,
	}, nil
}

// ListSomething ...
func (s *Server) ListSomething(ctx context.Context, req *v1.ListSomethingRequest) (*v1.ListSomethingResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	res := make([]*v1.Something, 0, len(s.data))
	for k, v := range s.data {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("json has invalid format, err: %v", err))
		}
		res = append(res, &v1.Something{
			Id:          k,
			Description: string(b),
		})
	}
	return &v1.ListSomethingResponse{
		Api:        apiVersion,
		Somethings: res,
	}, nil
}

// Login ...
func (s *Server) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}
	curUser := &v1.User{}
	for i := range s.userData {
		if s.userData[i].Username == req.User.Username {
			curUser = &s.userData[i]
			break
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(curUser.PasswordHash), []byte(req.User.Password))
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("failed to login, err: %v", err))
	}
	token, err := authutil.SignJWT(&v1.User{
		Id:        curUser.Id,
		CreatedAt: curUser.CreatedAt,
		Username:  curUser.Username,
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("failed to login, err: %v", err))
	}
	return &v1.LoginResponse{
		Api:   apiVersion,
		Token: token,
	}, nil
}

// Me ...
func (s *Server) Me(ctx context.Context, req *v1.MeRequest) (*v1.MeResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unknown, "metadata is required")
	}
	auth := md.Get("authorization")
	if len(auth[0]) == 0 {
		return nil, status.Error(codes.InvalidArgument, "authorization is required")
	}
	_, claims, err := authutil.ValidateJWT(auth[0])
	if err != nil {
		log.Print(err.Error())
	}
	return &v1.MeResponse{
		Api: apiVersion,
		User: &v1.User{
			Id:       claims["id"].(string),
			Username: claims["username"].(string),
		},
	}, nil
}

func (s *Server) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}
