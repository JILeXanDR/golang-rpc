package service

import (
	"github.com/JILeXanDR/golang-rpc/rpc/errors"
)

type Service struct {
}

func (s *Service) WithError(code errors.ErrorCode, message string) error {
	return errors.New(code, message)
}
