package encode

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateUserResponse struct{
		UserId string `json:"user_id"`
	}
	GetUserResponse struct{
		Email string `json:"email"`
	}
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
