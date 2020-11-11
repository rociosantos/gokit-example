package transport

import (
	"context"
	"net/http"

	"github.com/rociosantos/gokit-example/decode"
	"github.com/rociosantos/gokit-example/encode"
	"github.com/rociosantos/gokit-example/endpoint"

	"github.com/gorilla/mux"
	gokittransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, e endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/user").Handler(gokittransport.NewServer(
		e.CreateUser,
		decode.DecodeUserReq,
		encode.EncodeResponse,
	))
	r.Methods("GET").Path("/user/{id}").Handler(gokittransport.NewServer(
		e.GetUser,
		decode.DecodeEmailReq,
		encode.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}