package auth

import (
	"github.com/JILeXanDR/golang-rpc/pkg/jwt"
	"github.com/JILeXanDR/golang-rpc/rpc/errors"
	"net/http"
)

type LoginArgs struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginReply struct {
	TokenResponse
}

func (s *Service) Login(r *http.Request, args *LoginArgs, reply *LoginReply) error {
	if args.Login == "test" && args.Password == "test" {
		token, err := jwt.Sign()
		if err != nil {
			return err
		}
		reply.Token = token
		return nil
	}
	return errors.ErrBadLoginOrPassword
}
