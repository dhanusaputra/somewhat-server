package v1

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// Server ...
type Server struct {
	data map[string]interface{}
}

// NewServer ...
func NewServer(data map[string]interface{}) *Server {
	return &Server{data: data}
}

// GetSomething ...
func (s *Server) GetSomething(ctx context.Context, req *v1.GetSomethingRequest) (*v1.GetSomethingResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Id]; !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to find %v", req.Id))
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
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	return &v1.CreateSomethingResponse{
		Api: apiVersion,
		Id:  "",
	}, nil
}

// UpdateSomething ...
func (s *Server) UpdateSomething(ctx context.Context, req *v1.UpdateSomethingRequest) (*v1.UpdateSomethingResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	return &v1.UpdateSomethingResponse{
		Api:     apiVersion,
		Updated: true,
	}, nil
}

// DeleteSomething ...
func (s *Server) DeleteSomething(ctx context.Context, req *v1.DeleteSomethingRequest) (*v1.DeleteSomethingResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if _, ok := s.data[req.Id]; !ok {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to find %v", req.Id))
	}
	delete(s.data, req.Id)
	return &v1.DeleteSomethingResponse{
		Api:     apiVersion,
		Deleted: true,
	}, nil
}

// ListSomething ...
func (s *Server) ListSomething(ctx context.Context, req *v1.ListSomethingRequest) (*v1.ListSomethingResponse, error) {
	// check if the API version requested by client is supported by server
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

// checkAPI checks if the API version requested by client is supported by server
func (s *Server) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}
