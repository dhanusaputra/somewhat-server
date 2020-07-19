package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserServer ...
type UserServer struct {
	data map[string]interface{}
}

// NewUserServer ...
func NewUserServer(data map[string]interface{}) *UserServer {
	return &UserServer{data: data}
}

// GetUser ...
func (s *UserServer) GetUser(ctx context.Context, req *v1.GetSomethingRequest) (*v1.GetSomethingResponse, error) {
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

// CreateUser ...
func (s *UserServer) CreateUser(ctx context.Context, req *v1.CreateSomethingRequest) (*v1.CreateSomethingResponse, error) {
	// check if the API version requested by client is supported by server
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

// UpdateUser ...
func (s *UserServer) UpdateUser(ctx context.Context, req *v1.UpdateSomethingRequest) (*v1.UpdateSomethingResponse, error) {
	// check if the API version requested by client is supported by server
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

// DeleteUser ...
func (s *UserServer) DeleteUser(ctx context.Context, req *v1.DeleteSomethingRequest) (*v1.DeleteSomethingResponse, error) {
	// check if the API version requested by client is supported by server
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

// ListUser ...
func (s *UserServer) ListUser(ctx context.Context, req *v1.ListSomethingRequest) (*v1.ListSomethingResponse, error) {
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
func (s *UserServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}
