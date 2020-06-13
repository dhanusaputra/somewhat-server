package v1

import (
	"context"
	"testing"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

func TestGetSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer()
	got, err := s.GetSomething(ctx, &v1.GetSomethingRequest{Api: "v1", Id: "1"})
	if err != nil {
		t.Errorf("should be no error")
	}
	if got == nil {
		t.Errorf("should be something")
	}
}

func TestUpdateSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer()
	got, err := s.UpdateSomething(ctx, &v1.UpdateSomethingRequest{Api: "v1", Something: &v1.Something{}})
	if err != nil {
		t.Errorf("should be no error")
	}
	if got == nil {
		t.Errorf("should be something")
	}
}

func TestCreateSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer()
	got, err := s.CreateSomething(ctx, &v1.CreateSomethingRequest{Api: "v1", Something: &v1.Something{}})
	if err != nil {
		t.Errorf("should be no error")
	}
	if got == nil {
		t.Errorf("should be something")
	}
}

func TestListSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer()
	got, err := s.ListSomething(ctx, &v1.ListSomethingRequest{Api: "v1"})
	if err != nil {
		t.Errorf("should be no error")
	}
	if got == nil {
		t.Errorf("should be something")
	}
}

func TestDeleteSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer()
	got, err := s.DeleteSomething(ctx, &v1.DeleteSomethingRequest{Api: "v1", Id: "1"})
	if err != nil {
		t.Errorf("should be no error")
	}
	if got == nil {
		t.Errorf("should be something")
	}
}
