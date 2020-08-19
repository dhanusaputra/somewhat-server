package v1

import (
	"context"
	"reflect"
	"testing"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/stretchr/testify/assert"
)

var (
	testData = map[string]interface{}{
		"id": map[string]interface{}{"mockId": "mock"},
		"en": map[string]interface{}{"mockEn": "mock"},
		"in": make(chan int),
	}

	testUserData = map[string]interface{}{}
)

func TestGetSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer(testData, testUserData)
	type args struct {
		ctx context.Context
		req *v1.GetSomethingRequest
	}
	tests := []struct {
		name    string
		s       v1.SomewhatServer
		args    args
		want    *v1.GetSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetSomethingRequest{
					Api: "v1",
					Id:  "id",
				},
			},
			want: &v1.GetSomethingResponse{
				Api: "v1",
				Something: &v1.Something{
					Id:          "id",
					Description: "{\"mockId\":\"mock\"}",
				},
			},
		},
		{
			name: "unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetSomethingRequest{
					Api: "v2",
					Id:  "id",
				},
			},
			wantErr: true,
		},
		{
			name: "find ID failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetSomethingRequest{
					Api: "v1",
					Id:  "ps",
				},
			},
			wantErr: true,
		},
		{
			name: "marshal failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetSomethingRequest{
					Api: "v1",
					Id:  "in",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSomething(t *testing.T) {
	ctx := context.Background()
	s := NewServer(testData, testUserData)
	type args struct {
		ctx context.Context
		req *v1.UpdateSomethingRequest
	}
	tests := []struct {
		name    string
		s       v1.SomewhatServer
		args    args
		want    *v1.UpdateSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.UpdateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "id",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			want: &v1.UpdateSomethingResponse{
				Api:     "v1",
				Updated: true,
			},
		},
		{
			name: "unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.UpdateSomethingRequest{
					Api: "v2",
					Something: &v1.Something{
						Id:          "id",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "find ID failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.UpdateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "ps",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unmarshal failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.UpdateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "id",
						Description: "invalid",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSomething(t *testing.T) {
	ctx := context.Background()
	copyTestData := make(map[string]interface{}, len(testData))
	for k, v := range testData {
		copyTestData[k] = v
	}
	type args struct {
		ctx context.Context
		req *v1.CreateSomethingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.CreateSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &v1.CreateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "ps",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			want: &v1.CreateSomethingResponse{
				Api: "v1",
				Id:  "ps",
			},
		},
		{
			name: "unsupported API",
			args: args{
				ctx: ctx,
				req: &v1.CreateSomethingRequest{
					Api: "v2",
					Something: &v1.Something{
						Id:          "ps",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ID empty",
			args: args{
				ctx: ctx,
				req: &v1.CreateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ID already exits",
			args: args{
				ctx: ctx,
				req: &v1.CreateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "id",
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unmarshal failed",
			args: args{
				ctx: ctx,
				req: &v1.CreateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Id:          "ps",
						Description: "invalid",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer(testData, testUserData)
			defer func() {
				testData = copyTestData
			}()
			got, err := s.CreateSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListSomething(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *v1.ListSomethingRequest
	}
	tests := []struct {
		name    string
		args    args
		data    map[string]interface{}
		want    *v1.ListSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &v1.ListSomethingRequest{
					Api: "v1",
				},
			},
			data: map[string]interface{}{
				"id": map[string]interface{}{"mockId": "mock"},
				"en": map[string]interface{}{"mockEn": "mock"},
			},
			want: &v1.ListSomethingResponse{
				Api: "v1",
				Somethings: []*v1.Something{
					{
						Id:          "id",
						Description: "{\"mockId\":\"mock\"}",
					},
					{
						Id:          "en",
						Description: "{\"mockEn\":\"mock\"}",
					},
				},
			},
		},
		{
			name: "unsupported API",
			args: args{
				ctx: ctx,
				req: &v1.ListSomethingRequest{
					Api: "v2",
				},
			},
			wantErr: true,
		},
		{
			name: "marshal failed",
			data: map[string]interface{}{
				"id": map[string]interface{}{"mockId": "mock"},
				"en": map[string]interface{}{"mockEn": "mock"},
				"in": make(chan int),
			},
			args: args{
				ctx: ctx,
				req: &v1.ListSomethingRequest{
					Api: "v1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer(tt.data, testUserData)
			got, err := s.ListSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeleteSomething(t *testing.T) {
	ctx := context.Background()
	copyTestData := make(map[string]interface{}, len(testData))
	for k, v := range testData {
		copyTestData[k] = v
	}
	type args struct {
		ctx context.Context
		req *v1.DeleteSomethingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.DeleteSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &v1.DeleteSomethingRequest{
					Api: "v1",
					Id:  "id",
				},
			},
			want: &v1.DeleteSomethingResponse{
				Api:     "v1",
				Deleted: true,
			},
		},
		{
			name: "unsupported API",
			args: args{
				ctx: ctx,
				req: &v1.DeleteSomethingRequest{
					Api: "v2",
					Id:  "id",
				},
			},
			wantErr: true,
		},
		{
			name: "find ID failed",
			args: args{
				ctx: ctx,
				req: &v1.DeleteSomethingRequest{
					Api: "v1",
					Id:  "ps",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer(testData, testUserData)
			defer func() {
				testData = copyTestData
			}()
			got, err := s.DeleteSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
