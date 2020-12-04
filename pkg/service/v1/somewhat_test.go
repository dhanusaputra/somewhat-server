package v1

import (
	"context"
	"errors"
	"reflect"
	"testing"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/dhanusaputra/somewhat-server/util/authutil"
	"github.com/dhanusaputra/somewhat-server/util/testutil"
	"github.com/go-playground/validator"
	"github.com/mohae/deepcopy"
	"google.golang.org/grpc/metadata"
)

var (
	testData = map[string]interface{}{
		"id": map[string]interface{}{"mockId": "mock"},
		"en": map[string]interface{}{"mockEn": "mock"},
		"in": make(chan int),
	}

	testUserData = []v1.User{
		v1.User{
			Username:     "username",
			PasswordHash: "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO",
		},
	}
)

func TestGetSomething(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *v1.GetSomethingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.GetSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
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
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), nil)
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
	type args struct {
		ctx context.Context
		req *v1.UpdateSomethingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.UpdateSomethingResponse
		wantErr bool
	}{
		{
			name: "happy path",
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
			name: "ID empty",
			args: args{
				ctx: ctx,
				req: &v1.UpdateSomethingRequest{
					Api: "v1",
					Something: &v1.Something{
						Description: "{\"nothing\":\"nothing personal here\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "find ID failed",
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
			v := validator.New()
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), v)
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
						Id:          "invalid",
						Description: "invalid",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), v)
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
			s := NewServer(tt.data, deepcopy.Copy(testUserData).([]v1.User), nil)
			got, err := s.ListSomething(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSomething(t *testing.T) {
	ctx := context.Background()
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
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), nil)
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

func TestLogin(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req *v1.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.LoginResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Api: "v1",
					User: &v1.User{
						Username: "username",
						Password: "password",
					},
				},
			},
			want: &v1.LoginResponse{
				Api:   "v1",
				Token: "mockToken",
			},
			mock: func() {
				authutil.SignJWT = func(user *v1.User) (string, error) {
					return "mockToken", nil
				}
			},
		},
		{
			name: "user nill",
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Api:  "v1",
					User: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "unsupported API",
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Api:  "v2",
					User: &v1.User{},
				},
			},
			wantErr: true,
		},
		{
			name: "bcrypt failed",
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Api: "v1",
					User: &v1.User{
						Username: "username",
						Password: "password2",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "signJWT failed",
			args: args{
				ctx: ctx,
				req: &v1.LoginRequest{
					Api: "v1",
					User: &v1.User{
						Username: "username",
						Password: "password",
					},
				},
			},
			wantErr: true,
			mock: func() {
				authutil.SignJWT = func(user *v1.User) (string, error) {
					return "", errors.New("err")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT}).Restore()
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), nil)
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMe(t *testing.T) {
	md := metadata.New(map[string]string{"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjpudWxsLCJpZCI6IjEiLCJpc3MiOiJzb21ldGhpbmciLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.laPjiS5zWxCaihlGzYTI9jJ1lGuTWsTd4IJdEMgZwuc"})
	ctx := metadata.NewIncomingContext(context.Background(), md)
	md = metadata.New(map[string]string{"authorization": ""})
	ctxEmpty := metadata.NewIncomingContext(context.Background(), md)
	type args struct {
		ctx context.Context
		req *v1.MeRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.MeResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				req: &v1.MeRequest{
					Api: "v1",
				},
			},
			want: &v1.MeResponse{
				Api: "v1",
				User: &v1.User{
					Id:       "1",
					Username: "username",
				},
			},
		},
		{
			name: "unsupported API",
			args: args{
				ctx: ctx,
				req: &v1.MeRequest{
					Api: "v2",
				},
			},
			wantErr: true,
		},
		{
			name: "metadata empty",
			args: args{
				ctx: context.Background(),
				req: &v1.MeRequest{
					Api: "v1",
				},
			},
			wantErr: true,
		},
		{
			name: "authorization empty",
			args: args{
				ctx: ctxEmpty,
				req: &v1.MeRequest{
					Api: "v1",
				},
			},
			wantErr: true,
		},
		// {
		// 	name: "jwt invalid",
		// 	args: args{
		// 		ctx: ctx,
		// 		req: &v1.MeRequest{
		// 			Api: "v1",
		// 		},
		// 	},
		// 	want: &v1.MeResponse{
		// 		Api: "v1",
		// 		User: &v1.User{
		// 			Id:       "1",
		// 			Username: "username",
		// 		},
		// 	},
		// 	mock: func() {
		// 		authutil.ValidateJWT = func(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
		// 			return nil, jwt.MapClaims{
		// 				"created_at": nil,
		// 				"id":         "1",
		// 				"iss":        "something",
		// 				"username":   "username",
		// 			}, errors.New("err")
		// 		}
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.ValidateJWT}).Restore()
			s := NewServer(deepcopy.Copy(testData).(map[string]interface{}), deepcopy.Copy(testUserData).([]v1.User), nil)
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Me(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Me() = %v, want %v", got, tt.want)
			}
		})
	}
}
