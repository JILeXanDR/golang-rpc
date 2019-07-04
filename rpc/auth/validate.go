package auth

import (
	"github.com/JILeXanDR/golang-rpc/pkg/jwt"
	dgrijalva "github.com/dgrijalva/jwt-go"
	"net/http"
)

type ValidateReply struct {
	Status bool  `json:"status"`
	Err    error `json:"error"`
}

func (s *Service) Validate(r *http.Request, _, reply *ValidateReply) error {
	bearer := r.Header.Get("Authorization")
	err := jwt.Validate(bearer)
	if err != nil {
		validationErr := err.(*dgrijalva.ValidationError)
		reply.Status = false
		reply.Err = validationErr
	} else {
		reply.Status = true
	}
	return nil
}
