package auth

import (
	"errors"
	"net/http"
)

type LogoutArgs struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogoutReply struct {
	Token string `json:"token"`
}

func (s *Service) Logout(r *http.Request, args *LoginArgs, reply *LoginReply) error {
	return errors.New("not implemented")
}
