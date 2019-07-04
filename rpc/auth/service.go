package auth

import "github.com/JILeXanDR/golang-rpc/rpc/service"

type Service struct {
	service.Service
}

func New() *Service {
	return &Service{}
}
