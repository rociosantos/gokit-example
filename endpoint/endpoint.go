package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rociosantos/gokit-example/decode"
	"github.com/rociosantos/gokit-example/encode"
	"github.com/rociosantos/gokit-example/service"
)


type Endpoints struct{
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints {
		CreateUser: makeCreateUserEndpoint(s),
		GetUser: makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(decode.CreateUserRequest)
		userID, err := s.CreateUser(ctx, req.Email, req.Password)
		return encode.CreateUserResponse{UserId: userID}, err
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(decode.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)
		return encode.GetUserResponse{
			Email: email,
		}, err
	}
}