package main

import (
	"net/http"
	"log"
	"rpc/methods"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

func main() {
	s := rpc.NewServer()

	s.RegisterBeforeFunc(func(i *rpc.RequestInfo) {
		log.Printf("[hook before method call]: %v", i.Method)
	})

	s.RegisterAfterFunc(func(i *rpc.RequestInfo) {
		log.Printf("[hook after method call]: %v", i.Method)
		log.Println(i.Error)
	})

	s.RegisterInterceptFunc(func(i *rpc.RequestInfo) *http.Request {
		return i.Request
	})

	s.RegisterCodec(json.NewCodec(), "application/json")

	s.RegisterService(new(methods.HelloService), "Hello")
	s.RegisterService(new(methods.AuthorizeService), "Authorize")
	s.RegisterService(new(methods.PingService), "Ping")

	http.Handle("/rpc", s)

	debugMethods(s)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
