package main

import (
	server "github.com/JILeXanDR/golang-rpc/rpc"
	"github.com/JILeXanDR/golang-rpc/rpc/auth"
	"github.com/JILeXanDR/golang-rpc/rpc/errors"
	"github.com/JILeXanDR/golang-rpc/rpc/health"
	"github.com/gorilla/rpc/v2"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

func main() {

	v9 := validator.New()

	handler := server.NewRPCServer(
		map[string]interface{}{
			"Auth":   auth.New(),
			"Health": health.New(),
		},
	)

	handler.RegisterBeforeFunc(func(r *rpc.RequestInfo) {
		log.Printf("[RPC:BEFORE] %v", r.Method)
	})

	handler.RegisterAfterFunc(func(r *rpc.RequestInfo) {
		log.Printf("[RPC:AFTER] %v", r.Method)
	})

	handler.RegisterInterceptFunc(func(r *rpc.RequestInfo) *http.Request {
		log.Printf("[RPC:INTERCEPT] %v", r.Method)
		return r.Request
	})

	handler.RegisterValidateRequestFunc(func(r *rpc.RequestInfo, i interface{}) error {
		log.Printf("[RPC:VALIDATE] %v", r.Method)
		if err := v9.Struct(i); err != nil {
			return errors.NewValidationErr(err.(validator.ValidationErrors))
		}
		return nil
	})

	http.Handle("/rpc", handler)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
