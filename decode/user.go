package decode

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	GetUserRequest struct {
		Id string `json:"id"`
	}
)

func DecodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}
