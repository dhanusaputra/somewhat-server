package v1

import (
	"context"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

// Server ...
type Server struct{}

// NewServer ...
func NewServer() *Server {
	return &Server{}
}

// GetSomething ...
func (s *Server) GetSomething(ctx context.Context, req *v1.GetSomethingRequest) (*v1.GetSomethingResponse, error) {
	return &v1.GetSomethingResponse{
		Api:       apiVersion,
		Something: &v1.Something{},
	}, nil
}

// CreateSomething ...
func (s *Server) CreateSomething(ctx context.Context, req *v1.CreateSomethingRequest) (*v1.CreateSomethingResponse, error) {
	return &v1.CreateSomethingResponse{
		Api: apiVersion,
		Id:  "",
	}, nil
}

// UpdateSomething ...
func (s *Server) UpdateSomething(ctx context.Context, req *v1.UpdateSomethingRequest) (*v1.UpdateSomethingResponse, error) {
	return &v1.UpdateSomethingResponse{
		Api:     apiVersion,
		Updated: true,
	}, nil
}

// DeleteSomething ...
func (s *Server) DeleteSomething(ctx context.Context, req *v1.DeleteSomethingRequest) (*v1.DeleteSomethingResponse, error) {
	return &v1.DeleteSomethingResponse{
		Api:     apiVersion,
		Deleted: true,
	}, nil
}

// ListSomething ...
func (s *Server) ListSomething(ctx context.Context, req *v1.ListSomethingRequest) (*v1.ListSomethingResponse, error) {
	list := []*v1.Something{}
	return &v1.ListSomethingResponse{
		Api:        apiVersion,
		Somethings: list,
	}, nil
}
