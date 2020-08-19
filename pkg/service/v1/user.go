package v1

import (
	"context"
	"fmt"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser ...
func (s *Server) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Id]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("failed to find ID: %v", req.Id))
	}
	return &v1.GetUserResponse{
		Api: apiVersion,
	}, nil
}

// ListUser ...
func (s *Server) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	res := make([]string, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return &v1.ListUserResponse{
		Api: apiVersion,
	}, nil
}
